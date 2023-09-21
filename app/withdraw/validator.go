// desc:
// @author renshiwei
// Date: 2023/4/10 15:02

package withdraw

import (
	"context"
	"github.com/NodeDAO/oracle-go/common/errs"
	"github.com/NodeDAO/oracle-go/common/logger"
	"github.com/NodeDAO/oracle-go/consensus"
	"github.com/NodeDAO/oracle-go/contracts"
	"github.com/NodeDAO/oracle-go/contracts/withdrawOracle"
	"github.com/NodeDAO/oracle-go/eth1"
	consensusApi "github.com/attestantio/go-eth2-client/api/v1"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
	"math/big"
	"sort"
	"strconv"
)

// 1. Get all tokenIds
// 2. Query information about beacon
// 3. The pubkey not found by the beacon, balance = 32 GWEI
func (v *WithdrawHelper) obtainValidatorConsensusInfo(ctx context.Context) error {

	// Gets all active validators for the NodeDAO's StakingPool (nETH's vNFT)
	validatorBytesOfStakingPool, err := contracts.VnftContract.Contract.ActiveValidatorsOfStakingPool(nil)
	if err != nil {
		return errors.Wrap(err, "Failed to get VnftContract.ActiveValidatorsOfStakingPool")
	}
	if err := v.parseValidatorExaMap(ctx, validatorBytesOfStakingPool, LiquidStaking); err != nil {
		return errors.Wrap(err, "ActiveValidatorsOfStakingPool parseValidatorExaMap err.")
	}

	// Gets all active validators for the NodeDAO's User (vNFT)
	validatorBytesOfUser, err := contracts.VnftContract.Contract.ActiveValidatorOfUser(nil)
	if err != nil {
		return errors.Wrap(err, "Failed to get VnftContract.ActiveValidatorOfUser")
	}
	if err := v.parseValidatorExaMap(ctx, validatorBytesOfUser, USER); err != nil {
		return errors.Wrap(err, "ActiveValidatorOfUser parseValidatorExaMap err.")
	}

	return nil
}

func (v *WithdrawHelper) parseValidatorExaMap(ctx context.Context, validatorBytes [][]byte, vnftOwner VnftOwner) error {
	validatorCount := len(validatorBytes)

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
			v.ValidatorUnknownCount++
		}

		validatorExa := &ValidatorExa{
			Validator: validator,
			VnftOwner: vnftOwner,
		}
		v.validatorExaMap[pubkey] = validatorExa
	}

	for k, value := range validators {
		validatorExa := &ValidatorExa{
			Validator: value,
			VnftOwner: vnftOwner,
		}
		v.validatorExaMap[k] = validatorExa
	}

	return nil
}

func (v *WithdrawHelper) calculationValidatorExa(ctx context.Context) error {

	exitTokenIds := make([]*big.Int, 0)

	validatorExaArr := make([]*ValidatorExa, 0)

	for pubkey, exa := range v.validatorExaMap {
		// sum cl balance
		if exa.VnftOwner == LiquidStaking {
			v.clBalance = new(big.Int).Add(v.clBalance, big.NewInt(int64(exa.Validator.Balance)))
		}

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

		validatorExaArr = append(validatorExaArr, exa)
	}

	// cl balance to wei
	v.clBalance = eth1.GWEIToWEI(v.clBalance)

	sort.Slice(validatorExaArr, func(i, j int) bool {
		return validatorExaArr[i].TokenId.Uint64() < validatorExaArr[j].TokenId.Uint64()
	})

	// IsNeedOracleReportExit
	oracleReportExitNumbers, err := contracts.VnftContract.Contract.GetNftExitBlockNumbers(nil, exitTokenIds)
	if err != nil {
		return errors.Wrap(err, "Failed to get vnft'pubkey batch GetNftExitBlockNumbers")
	}

	for i, number := range oracleReportExitNumbers {
		if number.Cmp(decimal.Zero.BigInt()) == 0 {
			for _, exa := range validatorExaArr {
				pubkey := exa.Validator.Validator.PublicKey.String()
				if exitTokenIds[i] == exa.TokenId {
					exa.IsNeedOracleReportExit = true

					// slashed
					if exa.Validator.Status == consensusApi.ValidatorStateWithdrawalDone && exa.Validator.Validator.Slashed {
						validatorSlashedAmount, err := consensus.ConsensusClient.CustomizeBeaconService.ValidatorSlashedAmount(ctx, exa.Validator)
						if err != nil {
							return errors.Wrapf(err, "failed to get ValidatorSlashedAmount:%s", pubkey)
						}
						if validatorSlashedAmount.Cmp(eth1.ETH(2)) == 1 {
							logger.Warn("[WithdrawOracle] slash amount > 2 ETH.",
								zap.String("pubkey", pubkey),
								zap.String("slashAmount", validatorSlashedAmount.String()),
							)
							return errs.NewSleepError("slash amount > 2 ETH", RandomSleepTime())
						}

						exa.SlashAmount = validatorSlashedAmount
					} else {
						exa.SlashAmount = big.NewInt(0)
					}

					// ExitedAmountForClCapital
					if exa.VnftOwner == LiquidStaking {
						pubkeys := make([]string, 1)
						pubkeys[0] = pubkey
						validator, err := consensus.ConsensusClient.CustomizeBeaconService.ValidatorsByPubKey(ctx, exa.ExitedSlot.String(), pubkeys)
						if err != nil {
							return errors.Wrapf(err, "Failed to get validator ExitedAmountForClCapital. tokenId:%s pubkey:%s slot:%s", exa.TokenId.String(), pubkey, exa.ExitedSlot.String())
						}
						exa.ExitedAmountForClCapital = eth1.GWEIToWEI(big.NewInt(int64(validator[pubkey].Balance)))
					}

					// requireReportValidator If the limit is reached, it is not added
					exitLimit := int(v.exitRequestLimit.Int64())
					if len(v.requireReportValidator) < exitLimit {
						v.requireReportValidator[pubkey] = exa
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
