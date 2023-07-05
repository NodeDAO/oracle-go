// desc:
// @author renshiwei
// Date: 2023/4/10 15:02

package withdraw

import (
	"context"
	"github.com/NodeDAO/oracle-go/common/errs"
	"github.com/NodeDAO/oracle-go/consensus"
	"github.com/NodeDAO/oracle-go/contracts"
	"github.com/NodeDAO/oracle-go/contracts/withdrawOracle"
	"github.com/NodeDAO/oracle-go/eth1"
	consensusApi "github.com/attestantio/go-eth2-client/api/v1"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/params"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"math/big"
	"strconv"
)

// 1. Get all tokenIds
// 2. Query information about beacon
// 3. The pubkey not found by the beacon, balance = 32 GWEI
func (v *WithdrawHelper) obtainValidatorConsensusInfo(ctx context.Context) error {
	// Gets all active validators for the NodeDAO
	validatorBytes, err := contracts.VnftContract.Contract.ActiveValidatorsOfStakingPool(nil)
	if err != nil {
		return errors.Wrap(err, "Failed to get VnftContract.ActiveValidatorsOfStakingPool")
	}

	validatorCount := len(validatorBytes)
	v.validatorExaMap = make(map[string]*ValidatorExa)
	v.requireReportValidator = make(map[string]*ValidatorExa)

	pubkeys := make([]string, validatorCount)
	for i := 0; i < validatorCount; i++ {
		pubkey := hexutil.Encode(validatorBytes[i])
		pubkeys[i] = pubkey
	}

	validators, err := consensus.ConsensusClient.CustomizeBeaconService.ValidatorsByPubKey(ctx, v.refSlot.String(), pubkeys)
	if err != nil {
		return errors.Wrap(err, "Failed to get validators info.")
	}
	for _, pubkey := range pubkeys {
		validator, ok := validators[pubkey]
		// Handle pubkeys that are not query by Beacon
		// set balance = 32 GWEI
		if !ok {
			validator = &consensusApi.Validator{
				Balance: 32e9,
			}
		}

		validatorExa := &ValidatorExa{
			Validator: validator,
		}
		v.validatorExaMap[pubkey] = validatorExa
	}

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
		pubkeyBytes, err := hexutil.Decode(pubkey)
		if err != nil {
			return errors.Wrapf(err, "failed to decode pubkey: %s", pubkey)
		}

		// get tokenId
		tokenId, err := contracts.VnftContract.Contract.TokenOfValidator(nil, pubkeyBytes)
		if err != nil {
			return errors.Wrapf(err, "Failed to get vnft'pubkey tokenId. pubkey:%s", pubkey)
		}
		exa.TokenId = tokenId

		// get OperatorId
		operatorId, err := contracts.VnftContract.Contract.OperatorOf(nil, tokenId)
		if err != nil {
			return errors.Wrapf(err, "Failed to get vnft'pubkey operatorId. pubkey:%s tokenId:%v", pubkey, tokenId)
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
			if executionBlock.BlockNumber.Cmp(big.NewInt(0)) == 0 {
				return errors.Errorf("Execution block is zero err.tokenId:%s pubkey:%s", exa.TokenId, pubkey)
			}
			exa.ExitedBlockHeight = executionBlock.BlockNumber

			// IsNeedOracleReportExit
			exitTokenIds = append(exitTokenIds, tokenId)
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
			// todo 遍历map 改为数组
			for s, exa := range v.validatorExaMap {
				if exitTokenIds[i] == exa.TokenId {
					exa.IsNeedOracleReportExit = true

					// slashed
					if exa.Validator.Validator.Slashed {
						withdrawAbleSlot := consensus.ConsensusClient.ChainTimeService.EpochToSlot(exa.Validator.Validator.WithdrawableEpoch)
						currentSlot := consensus.ConsensusClient.ChainTimeService.CurrentSlot()
						if withdrawAbleSlot <= currentSlot {
							pubkeys := make([]string, 1)
							pubkeys[0] = s
							validator, err := consensus.ConsensusClient.CustomizeBeaconService.ValidatorsByPubKey(ctx, strconv.FormatInt(int64(withdrawAbleSlot), 10), pubkeys)
							if err != nil {
								return errors.Wrapf(err, "Failed to get validator ExitedAmount. tokenId:%s pubkey:%s slot:%s", exa.TokenId.String(), s, exa.ExitedSlot.String())
							}
							exa.SlashAmount = eth1.ETH32().Sub(decimal.NewFromInt(int64(validator[s].Balance)).Mul(decimal.NewFromInt(params.GWei))).BigInt()
						}
					} else {
						exa.SlashAmount = big.NewInt(0)
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

					// requireReportValidator If the limit is reached, it is not added
					exitLimit := int(v.exitRequestLimit.Int64())
					if len(v.requireReportValidator) < exitLimit {
						v.requireReportValidator[s] = exa
					}

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

	if balance.Cmp(big.NewInt(0)) == 0 {
		return errs.NewSleepError("ClVaultBalance is zero. Cancel report.", RandomSleepTime())
	}

	v.clVaultBalance = balance

	return nil
}

func (v *WithdrawHelper) calculationExitValidatorInfo(ctx context.Context) error {
	for _, exa := range v.requireReportValidator {
		if exa.ExitedBlockHeight.Cmp(big.NewInt(0)) == 0 {
			continue
		}
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
