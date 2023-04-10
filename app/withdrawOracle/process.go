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
	v.withdrawInfos = make([]*withdrawOracle.WithdrawInfo, 0)
	v.exitValidatorInfos = make([]*withdrawOracle.ExitValidatorInfo, 0)

	return nil
}

// 1.
func (v *WithdrawHelper) obtainReportData(ctx context.Context) error {

	return nil
}
