// desc:
// @author renshiwei
// Date: 2023/4/7 15:33

package withdrawOracle

import (
	"context"
	"github.com/NodeDAO/oracle-go/consensus"
	"github.com/NodeDAO/oracle-go/contracts"
	"github.com/NodeDAO/oracle-go/eth1"
	consensusApi "github.com/attestantio/go-eth2-client/api/v1"
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

	// todo
	IsDelayedExit bool
}

type WithdrawHelper struct {
	RefSlot *big.Int

	ValidatorExaMap        map[string]*ValidatorExa
	RequireReportValidator map[string]*ValidatorExa

	ClBalance      *big.Int
	ClVaultBalance *big.Int
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

	for s, exa := range v.ValidatorExaMap {
		// get tokenId
		tokenId, err := contracts.VnftContract.TokenOfValidator(nil, []byte(s))
		if err != nil {
			return errors.Wrapf(err, "Failed to get vnft's tokenId. pubkey:%s", s)
		}
		exa.TokenId = tokenId

		// get OperatorId
		operatorId, err := contracts.VnftContract.OperatorOf(nil, tokenId)
		if err != nil {
			return errors.Wrapf(err, "Failed to get vnft's operatorId. pubkey:%s tokenId:%v", s, tokenId)
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

			}
			parseUint, _ := strconv.ParseInt(executionBlock.BlockNumber, 10, 64)
			exa.ExitedBlockHeight = big.NewInt(parseUint)

			// IsOracleReportExit
			exitTokenIds = append(exitTokenIds, tokenId)
		}

	}

	// IsOracleReportExit
	oracleReportExitNumbers, err := contracts.VnftContract.GetNftExitBlockNumbers(nil, exitTokenIds)
	if err != nil {
		return errors.Wrap(err, "Failed to get vnft's batch GetNftExitBlockNumbers")
	}

	for i, number := range oracleReportExitNumbers {
		if !decimal.NewFromBigInt(number, 0).LessThanOrEqual(decimal.Zero) {
			for s, exa := range v.ValidatorExaMap {
				if exitTokenIds[i] == exa.TokenId {
					exa.IsOracleReportExit = true

					// slashed
					if exa.Validator.Validator.Slashed {
						exa.SlashAmount = eth1.ETH32().Sub(decimal.NewFromInt(int64(exa.Validator.Balance)).Mul(decimal.NewFromInt(1e9))).BigInt()
					}

					// RequireReportValidator
					v.RequireReportValidator[s] = exa

					break
				}
			}
		}
	}

	// todo delayedExitSlashStandard

	return nil
}
