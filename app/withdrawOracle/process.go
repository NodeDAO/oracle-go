// desc:
// @author renshiwei
// Date: 2023/4/10 15:03

package withdrawOracle

import (
	"context"
	"github.com/NodeDAO/oracle-go/consensus"
	"github.com/NodeDAO/oracle-go/contracts"
	"github.com/NodeDAO/oracle-go/contracts/withdrawOracle"
	"github.com/pkg/errors"
	"math/big"
	"sort"
)

// core process
// 1. init: setup
// 2. get Validators
// 3. calculation ValidatorExa and delayedExitTokenIds
// 4. clVaultBalance
// 5. calculationExitValidatorInfo
// 6. calculationForOperator for rewards
// 7. dealLargeExitDelayedRequest for LargeExitDelayedRequestIds
// 8. obtainReportData
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

	if err := v.calculationExitValidatorInfo(ctx); err != nil {
		return errors.Unwrap(err)
	}

	if err := v.calculationForOperator(ctx); err != nil {
		return errors.Unwrap(err)
	}

	if err := v.dealLargeExitDelayedRequest(ctx); err != nil {
		return errors.Unwrap(err)
	}
	// todo
	if err := v.obtainReportData(ctx); err != nil {
		return errors.Unwrap(err)
	}

	return nil
}

func (v *WithdrawHelper) setup(ctx context.Context) error {
	// executionBlock
	executionBlock, err := consensus.ConsensusClient.CustomizeBeaconService.ExecutionBlock(ctx, v.RefSlot.String())
	if err != nil {
		return errors.Wrapf(err, "Failed to get execution block. slot:%s", v.RefSlot.String())
	}
	v.executionBlock = executionBlock

	// delayedExitSlashStandard
	v.delayedExitSlashStandard, err = contracts.LiqContract.Contract.DelayedExitSlashStandard(nil)
	if err != nil {
		return errors.Wrap(err, "Failed to get LiqContract delayedExitSlashStandard.")
	}

	// clVaultMinSettleLimit
	v.clVaultMinSettleLimit, err = contracts.WithdrawOracleContract.Contract.ClVaultMinSettleLimit(nil)
	if err != nil {
		return errors.Wrap(err, "Failed to get WithdrawOracleContract clVaultMinSettleLimit.")
	}

	v.totalNftCount, err = contracts.VnftContract.Contract.TotalSupply(nil)
	if err != nil {
		return errors.Wrap(err, "Failed to get VnftContract TotalSupply.")
	}

	v.delayedExitTokenIds = make([]*big.Int, 0)
	v.largeExitDelayedRequestIds = make([]*big.Int, 0)
	v.withdrawInfos = make([]withdrawOracle.WithdrawInfo, 0)
	v.exitValidatorInfos = make([]withdrawOracle.ExitValidatorInfo, 0)

	return nil
}

// Slice types need to be sorted
func (v *WithdrawHelper) obtainReportData(ctx context.Context) error {
	// sort slice

	sort.Slice(v.withdrawInfos, func(i, j int) bool {
		return v.withdrawInfos[i].OperatorId < v.withdrawInfos[j].OperatorId
	})

	sort.Slice(v.exitValidatorInfos, func(i, j int) bool {
		return v.exitValidatorInfos[i].ExitTokenId < v.exitValidatorInfos[j].ExitTokenId
	})

	sort.Slice(v.delayedExitTokenIds, func(i, j int) bool {
		return v.delayedExitTokenIds[i].Cmp(v.delayedExitTokenIds[j]) < 0
	})

	sort.Slice(v.largeExitDelayedRequestIds, func(i, j int) bool {
		return v.largeExitDelayedRequestIds[i].Cmp(v.largeExitDelayedRequestIds[j]) < 0
	})

	v.reportData = &withdrawOracle.WithdrawOracleReportData{
		ConsensusVersion:           v.ConsensusVersion,
		RefSlot:                    v.RefSlot,
		ClBalance:                  v.clBalance,
		ClVaultBalance:             v.clVaultBalance,
		ReportExitedCount:          big.NewInt(int64(len(v.exitValidatorInfos))),
		WithdrawInfos:              v.withdrawInfos,
		ExitValidatorInfos:         v.exitValidatorInfos,
		DelayedExitTokenIds:        v.delayedExitTokenIds,
		LargeExitDelayedRequestIds: v.largeExitDelayedRequestIds,
	}

	return nil
}
