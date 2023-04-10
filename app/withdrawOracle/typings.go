// desc:
// @author renshiwei
// Date: 2023/4/7 15:33

package withdrawOracle

import (
	"context"
	"github.com/NodeDAO/oracle-go/config"
	"github.com/NodeDAO/oracle-go/consensus"
	"github.com/NodeDAO/oracle-go/consensus/beacon"
	"github.com/NodeDAO/oracle-go/contracts"
	"github.com/NodeDAO/oracle-go/eth1"
	consensusApi "github.com/attestantio/go-eth2-client/api/v1"
	"github.com/ethereum/go-ethereum/params"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"math/big"
	"strconv"
)

type ValidatorExa struct {
	Validator *consensusApi.Validator

	IsExited          bool
	ExitedSlot        *big.Int
	ExitedBlockHeight *big.Int

	TokenId            *big.Int
	OperatorId         *big.Int
	IsOracleReportExit bool

	// 1.slashed 2.exited 3.Not OracleReportExit
	SlashAmount *big.Int

	IsDelayedExit bool
}

type WithdrawHelper struct {
	RefSlot                  *big.Int
	ExecutionBlock           *beacon.ExecutionBlock
	DelayedExitSlashStandard *big.Int

	ValidatorExaMap        map[string]*ValidatorExa
	RequireReportValidator map[string]*ValidatorExa

	ClBalance      *big.Int
	ClVaultBalance *big.Int
}

// core process
// 1. init: setup
// 2. get Validators
// 3. calculation ValidatorExa
// 4. ClVaultBalance
// 5.
func (v *WithdrawHelper) process(ctx context.Context) error {
	if err := v.setup(ctx); err != nil {
		return errors.Unwrap(err)
	}

	if err := v.obtainValidatorConsensusInfo(ctx); err != nil {
		return errors.Unwrap(err)
	}

	if err := v.calculationValidatorExa(ctx); err != nil {
		return errors.Unwrap(err)
	}

	if err := v.calculationClVaultBalance(ctx); err != nil {
		return errors.Unwrap(err)
	}

	return nil
}

func (v *WithdrawHelper) setup(ctx context.Context) error {
	// ExecutionBlock
	executionBlock, err := consensus.ConsensusClient.CustomizeBeaconService.ExecutionBlock(ctx, v.RefSlot.String())
	if err != nil {
		return errors.Wrapf(err, "Failed to get execution block. slot:%s", v.RefSlot.String())
	}
	v.ExecutionBlock = executionBlock

	// DelayedExitSlashStandard
	v.DelayedExitSlashStandard, err = contracts.LiqContract.DelayedExitSlashStandard(nil)
	if err != nil {
		return errors.Wrap(err, "Failed to get LiqContract DelayedExitSlashStandard.")
	}

	return nil
}

func (v *WithdrawHelper) obtainValidatorConsensusInfo(ctx context.Context) error {
	// Gets all active validators for the NodeDAO
	validatorBytes, err := contracts.VnftContract.ActiveValidatorsOfStakingPool(nil)
	if err != nil {
		return errors.Wrap(err, "Failed to get VnftContract.ActiveValidatorsOfStakingPool")
	}

	validatorCount := len(validatorBytes)
	v.ValidatorExaMap = make(map[string]*ValidatorExa, 0)
	v.RequireReportValidator = make(map[string]*ValidatorExa, 0)

	pubkeys := make([]string, validatorCount)
	for i := 0; i < validatorCount; i++ {
		pubkeys[i] = string(validatorBytes[i])
	}

	validators, err := consensus.ConsensusClient.CustomizeBeaconService.ValidatorsByPubKey(ctx, v.RefSlot.String(), pubkeys)
	for k, value := range validators {
		validatorExa := &ValidatorExa{
			Validator: value,
		}
		v.ValidatorExaMap[k] = validatorExa
	}

	return nil
}

func (v *WithdrawHelper) calculationValidatorExa(ctx context.Context) error {

	exitTokenIds := make([]*big.Int, 0)

	for pubkey, exa := range v.ValidatorExaMap {
		// sum cl balance
		v.ClBalance = new(big.Int).Add(v.ClBalance, big.NewInt(int64(exa.Validator.Balance)))

		// get tokenId
		tokenId, err := contracts.VnftContract.TokenOfValidator(nil, []byte(pubkey))
		if err != nil {
			return errors.Wrapf(err, "Failed to get vnft'pubkey tokenId. pubkey:%pubkey", pubkey)
		}
		exa.TokenId = tokenId

		// get OperatorId
		operatorId, err := contracts.VnftContract.OperatorOf(nil, tokenId)
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

			// IsOracleReportExit
			exitTokenIds = append(exitTokenIds, tokenId)
		}

		// not exited
		// delayedExitSlashStandard
		if !exa.IsExited {
			err := v.calculationIsDelayedExit(ctx, exa)
			if err != nil {
				if err != nil {
					return errors.Wrapf(err, "Failed to calculationIsDelayedExit. tokenId:%pubkey pubkey:%pubkey", exa.TokenId.String(), pubkey)
				}
			}
		}

	}

	// cl balance to wei
	v.ClBalance = eth1.GWEIToWEI(v.ClBalance)

	// IsOracleReportExit
	oracleReportExitNumbers, err := contracts.VnftContract.GetNftExitBlockNumbers(nil, exitTokenIds)
	if err != nil {
		return errors.Wrap(err, "Failed to get vnft'pubkey batch GetNftExitBlockNumbers")
	}

	for i, number := range oracleReportExitNumbers {
		// Ge 0
		if number.Cmp(decimal.Zero.BigInt()) == 1 {
			for s, exa := range v.ValidatorExaMap {
				if exitTokenIds[i] == exa.TokenId {
					exa.IsOracleReportExit = true

					// slashed
					if exa.Validator.Validator.Slashed {
						exa.SlashAmount = eth1.ETH32().Sub(decimal.NewFromInt(int64(exa.Validator.Balance)).Mul(decimal.NewFromInt(params.GWei))).BigInt()
					}

					// RequireReportValidator
					v.RequireReportValidator[s] = exa

					break
				}
			}
		}
	}

	return nil
}

// ClVaultBalance
func (v *WithdrawHelper) calculationClVaultBalance(ctx context.Context) error {
	executionBlock := v.ExecutionBlock

	balance, err := eth1.ElClient.BalanceAt(ctx, config.Config.Contracts.ClVault, executionBlock.BlockNumber)
	if err != nil {
		return errors.Wrapf(err, "Failed to get execution block. executionBlock:%s address:%s", executionBlock.BlockNumber.String(), config.Config.Contracts.ClVault)
	}
	v.ClVaultBalance = balance

	return nil
}

func (v *WithdrawHelper) calculationIsDelayedExit(ctx context.Context, exa *ValidatorExa) error {
	// View liq contract nftUnstakeBlockNumbers and query tokenId exit block height (block height is 0, exit not initiated)
	requireBlockNumbers, err := contracts.LiqContract.NftUnstakeBlockNumbers(nil, exa.TokenId)
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
		reportDelayedNumber, err := contracts.LiqContract.NftExitDelayedSlashRecords(nil, exa.TokenId)
		if err != nil {
			return errors.Wrapf(err, "Failed to get liq contract NftExitDelayedSlashRecords. tokenId:%s", exa.TokenId.String())
		}

		// If reportDelayedNumber == 0, the report is not reported. curNumber - requireBlockNumbers > (delayedExitSlashStandard = 21600)
		// If reportDelayedNumber is greater than 0, it has been reported and has not logged out in time. curNumber-reportDelayedNumber > (delayedExitSlashStandard = 21600)
		isReportDelayed := false
		curNumber := v.ExecutionBlock.BlockNumber
		if reportDelayedNumber.Cmp(decimal.Zero.BigInt()) == 0 {
			if new(big.Int).Sub(curNumber, requireBlockNumbers).Cmp(v.DelayedExitSlashStandard) == 1 {
				isReportDelayed = true
			}
		} else if reportDelayedNumber.Cmp(decimal.Zero.BigInt()) == 1 {
			if new(big.Int).Sub(curNumber, reportDelayedNumber).Cmp(v.DelayedExitSlashStandard) == 1 {
				isReportDelayed = true
			}
		}

		if isReportDelayed {
			exa.IsDelayedExit = true
		}

	}

	return nil
}
