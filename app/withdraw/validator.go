// desc:
// @author renshiwei
// Date: 2023/4/10 15:02

package withdraw

import (
	"context"
	"github.com/NodeDAO/oracle-go/consensus"
	"github.com/NodeDAO/oracle-go/contracts"
	"github.com/NodeDAO/oracle-go/contracts/withdrawOracle"
	"github.com/NodeDAO/oracle-go/eth1"
	consensusApi "github.com/attestantio/go-eth2-client/api/v1"
	"github.com/ethereum/go-ethereum/params"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"math/big"
	"strconv"
)

func (v *WithdrawHelper) obtainValidatorConsensusInfo(ctx context.Context) error {
	// Gets all active validators for the NodeDAO
	validatorBytes, err := contracts.VnftContract.Contract.ActiveValidatorsOfStakingPool(nil)
	if err != nil {
		return errors.Wrap(err, "Failed to get VnftContract.ActiveValidatorsOfStakingPool")
	}

	validatorCount := len(validatorBytes)
	v.validatorExaMap = make(map[string]*ValidatorExa, 0)
	v.requireReportValidator = make(map[string]*ValidatorExa, 0)

	pubkeys := make([]string, validatorCount)
	for i := 0; i < validatorCount; i++ {
		pubkeys[i] = string(validatorBytes[i])
	}

	validators, err := consensus.ConsensusClient.CustomizeBeaconService.ValidatorsByPubKey(ctx, v.refSlot.String(), pubkeys)
	for k, value := range validators {
		validatorExa := &ValidatorExa{
			Validator: value,
		}
		v.validatorExaMap[k] = validatorExa
	}

	return nil
}

func (v *WithdrawHelper) calculationValidatorExa(ctx context.Context) error {

	exitTokenIds := make([]*big.Int, 0)

	for pubkey, exa := range v.validatorExaMap {
		// sum cl balance
		v.clBalance = new(big.Int).Add(v.clBalance, big.NewInt(int64(exa.Validator.Balance)))

		// get tokenId
		tokenId, err := contracts.VnftContract.Contract.TokenOfValidator(nil, []byte(pubkey))
		if err != nil {
			return errors.Wrapf(err, "Failed to get vnft'pubkey tokenId. pubkey:%pubkey", pubkey)
		}
		exa.TokenId = tokenId

		// get OperatorId
		operatorId, err := contracts.VnftContract.Contract.OperatorOf(nil, tokenId)
		if err != nil {
			return errors.Wrapf(err, "Failed to get vnft'pubkey operatorId. pubkey:%pubkey tokenId:%v", pubkey, tokenId)
		}
		exa.OperatorId = operatorId

		// IsExited
		if exa.Validator.Status == consensusApi.ValidatorStateWithdrawalDone && exa.Validator.Balance == 0 {
			exa.IsExited = true
			// ExitedSlot
			exitedSlot := consensus.ConsensusClient.ChainTimeService.FirstSlotOfEpoch(exa.Validator.Validator.ExitEpoch)
			exa.ExitedSlot = big.NewInt(int64(exitedSlot))

			// ExitedBlockHeight
			executionBlock, err := consensus.ConsensusClient.CustomizeBeaconService.ExecutionBlock(ctx, strconv.Itoa(int(exitedSlot)))
			if err != nil {
				return errors.Wrapf(err, "Failed to get execution block. slot:%v", exitedSlot)
			}
			exa.ExitedBlockHeight = executionBlock.BlockNumber

			// IsNeedOracleReportExit
			exitTokenIds = append(exitTokenIds, tokenId)
		}

		// not exited
		// delayedExitSlashStandard
		if !exa.IsExited {
			err := v.calculationIsDelayedExit(ctx, exa)
			if err != nil {
				if err != nil {
					return errors.Wrapf(err, "Failed to calculationIsDelayedExit. tokenId:%s  pubkey:%s", exa.TokenId.String(), pubkey)
				}
			}
		}

	}

	// cl balance to wei
	v.clBalance = eth1.GWEIToWEI(v.clBalance)

	// IsNeedOracleReportExit
	oracleReportExitNumbers, err := contracts.VnftContract.Contract.GetNftExitBlockNumbers(nil, exitTokenIds)
	if err != nil {
		return errors.Wrap(err, "Failed to get vnft'pubkey batch GetNftExitBlockNumbers")
	}

	for i, number := range oracleReportExitNumbers {
		// Ge 0
		if number.Cmp(decimal.Zero.BigInt()) == 0 {
			for s, exa := range v.validatorExaMap {
				if exitTokenIds[i] == exa.TokenId {
					exa.IsNeedOracleReportExit = true

					// slashed
					if exa.Validator.Validator.Slashed {
						exa.SlashAmount = eth1.ETH32().Sub(decimal.NewFromInt(int64(exa.Validator.Balance)).Mul(decimal.NewFromInt(params.GWei))).BigInt()
					}

					// ExitedAmount
					isOwnerLiqPool, err := v.IsOwnerLiqPool(ctx, exa)
					if err != nil {
						return errors.Wrap(err, "")
					}
					if isOwnerLiqPool {
						pubkeys := make([]string, 1)
						pubkeys[0] = s
						validator, err := consensus.ConsensusClient.CustomizeBeaconService.ValidatorsByPubKey(ctx, exa.ExitedSlot.String(), pubkeys)
						if err != nil {
							return errors.Wrapf(err, "Failed to get validator ExitedAmount. tokenId:%s pubkey:%s slot:%s", exa.TokenId.String(), s, exa.ExitedSlot.String())
						}
						exa.ExitedAmount = eth1.GWEIToWEI(big.NewInt(int64(validator[s].Balance)))
					}

					// requireReportValidator
					v.requireReportValidator[s] = exa
					break
				}
			}
		}
	}

	return nil
}

// clVaultBalance
func (v *WithdrawHelper) calculationClVaultBalance(ctx context.Context) error {
	executionBlock := v.executionBlock

	balance, err := eth1.ElClient.BalanceAt(ctx, contracts.GetClVaultAddress(), executionBlock.BlockNumber)
	if err != nil {
		return errors.Wrapf(err, "Failed to get ClVault BalanceAt. executionBlock:%s address:%s", executionBlock.BlockNumber.String(), contracts.GetClVaultAddress())
	}
	v.clVaultBalance = balance

	return nil
}

func (v *WithdrawHelper) calculationIsDelayedExit(ctx context.Context, exa *ValidatorExa) error {
	// View liq contract nftUnstakeBlockNumbers and query tokenId exit block height (block height is 0, exit not initiated)
	requireBlockNumbers, err := contracts.LiqContract.Contract.NftUnstakeBlockNumbers(nil, exa.TokenId)
	if err != nil {
		return errors.Wrapf(err, "Failed to get liq contract nftUnstakeBlockNumbers. tokenId:%s", exa.TokenId.String())
	}

	// ge 0 requested to exit. need to exit
	if requireBlockNumbers.Cmp(decimal.Zero.BigInt()) == 1 {
		// delayedExitSlashRecords The number of blocks that are not exited in time is reported by the Oracle database.
		// This record is used to handle the report rounds that do not exit the validator in time
		//  1. If the block height is 0, the oracle does not report the token Id
		//  2. Determine the current block height - The last reported block height > The maximum block height that can not exit in time
		// (delayedExitSlashStandard = 21600)
		reportDelayedNumber, err := contracts.LiqContract.Contract.NftExitDelayedSlashRecords(nil, exa.TokenId)
		if err != nil {
			return errors.Wrapf(err, "Failed to get liq contract NftExitDelayedSlashRecords. tokenId:%s", exa.TokenId.String())
		}

		// If reportDelayedNumber == 0, the report is not reported. curNumber - requireBlockNumbers > (delayedExitSlashStandard = 21600)
		// If reportDelayedNumber is greater than 0, it has been reported and has not logged out in time. curNumber-reportDelayedNumber > (delayedExitSlashStandard = 21600)
		isReportDelayed := false
		curNumber := v.executionBlock.BlockNumber
		if reportDelayedNumber.Cmp(decimal.Zero.BigInt()) == 0 {
			if new(big.Int).Sub(curNumber, requireBlockNumbers).Cmp(v.delayedExitSlashStandard) == 1 {
				isReportDelayed = true
			}
		} else if reportDelayedNumber.Cmp(decimal.Zero.BigInt()) == 1 {
			if new(big.Int).Sub(curNumber, reportDelayedNumber).Cmp(v.delayedExitSlashStandard) == 1 {
				isReportDelayed = true
			}
		}

		if isReportDelayed {
			exa.IsDelayedExit = true
			v.delayedExitTokenIds = append(v.delayedExitTokenIds, exa.TokenId)
		}

	}

	return nil
}

func (v *WithdrawHelper) calculationExitValidatorInfo(ctx context.Context) error {
	for _, exa := range v.requireReportValidator {
		w := withdrawOracle.ExitValidatorInfo{
			ExitTokenId:     exa.TokenId.Uint64(),
			ExitBlockNumber: exa.ExitedBlockHeight,
			SlashAmount:     exa.SlashAmount,
		}
		v.exitValidatorInfos = append(v.exitValidatorInfos, w)
	}
	return nil
}

func (v *WithdrawHelper) IsOwnerLiqPool(ctx context.Context, exa *ValidatorExa) (bool, error) {
	add, err := contracts.VnftContract.Contract.OwnerOf(nil, exa.TokenId)
	if err != nil {
		return false, errors.Wrapf(err, "Failed to get VnftContract OwnerOf. tokenId:%s", exa.TokenId.String())
	}

	return add.String() == contracts.LiqContract.Address, nil
}

// dealLargeExitDelayedRequest
func (v *WithdrawHelper) dealLargeExitDelayedRequest(ctx context.Context) error {
	// Get all the operators
	operatorCount, err := v.getAllOperatorId(ctx)
	if err != nil {
		return errors.Wrap(err, "Failed to get OperatorContract GetNodeOperatorsCount.")
	}

	// operator id start 0.
	for i := int64(1); i < operatorCount.Int64()+1; i++ {

		operatorId := big.NewInt(i)
		operatorPendingEthRequestAmount, err := contracts.LiqContract.Contract.OperatorPendingEthRequestAmount(nil, operatorId)
		if err != nil {
			return errors.Wrapf(err, "Failed to get LiqContract OperatorPendingEthRequestAmount. OperatorId: %s", operatorId.String())
		}

		operatorPendingEthPoolBalance, err := contracts.LiqContract.Contract.OperatorPendingEthPoolBalance(nil, operatorId)
		if err != nil {
			return errors.Wrapf(err, "Failed to get LiqContract operatorPendingEthPoolBalance. OperatorId: %s", operatorId.String())
		}

		// operatorPendingEthRequestAmount > operatorPendingEthPoolBalance
		// has largeExitDelayedRequest
		if operatorPendingEthRequestAmount.Cmp(operatorPendingEthPoolBalance) == 1 {
			cha := new(big.Int).Sub(operatorPendingEthRequestAmount, operatorPendingEthPoolBalance)

			withdrawalQueues, err := contracts.LiqContract.Contract.GetWithdrawalOfOperator(nil, operatorId)
			if err != nil {
				return errors.Wrapf(err, "Failed to get LiqContract getWithdrawalOfOperator. OperatorId: %s", operatorId.String())
			}

			for i := len(withdrawalQueues) - 1; i >= 0; i-- {
				q := withdrawalQueues[i]
				// block height
				isDelayed := new(big.Int).Sub(v.executionBlock.BlockNumber, q.WithdrawHeight).Cmp(v.delayedExitSlashStandard) == 1

				if isDelayed {
					v.largeExitDelayedRequestIds = append(v.largeExitDelayedRequestIds, big.NewInt(int64(i)))
				}

				if q.ClaimEthAmount.Cmp(cha) == 1 {
					break
				}
			}
		}
	}

	return nil
}
