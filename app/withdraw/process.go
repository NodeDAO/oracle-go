// desc:
// @author renshiwei
// Date: 2023/4/10 15:03

package withdraw

import (
	"context"
	"github.com/NodeDAO/oracle-go/app/consensusModule"
	"github.com/NodeDAO/oracle-go/common/logger"
	"github.com/NodeDAO/oracle-go/consensus"
	"github.com/NodeDAO/oracle-go/contracts"
	"github.com/NodeDAO/oracle-go/contracts/withdrawOracle"
	"github.com/NodeDAO/oracle-go/eth1"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"math/big"
	"sort"
	"time"
)

func (v *WithdrawHelper) ProcessReport(ctx context.Context) error {
	paused, err := v.oracle.Paused(ctx)
	if err != nil {
		return errors.Wrap(err, "")
	}
	if paused {
		logger.Info("withdrawOracle is paused.")
		time.Sleep(time.Second * SECONDS_PER_EPOCH)
	}

	if err := v.buildReportData(ctx); err != nil {
		return errors.Wrap(err, "")
	}

	reportDataHash, err := EncodeReportData(v.reportData)
	if err != nil {
		return errors.Wrap(err, "")
	}

	if err := v.hashConsensusHelper.ProcessReportHash(ctx, reportDataHash, v.refSlot, v.consensusVersion); err != nil {
		return errors.Wrap(err, "")
	}

	if err := v.processReportData(ctx, reportDataHash); err != nil {
		return errors.Wrap(err, "")
	}

	return nil
}

// buildReportData core process
// 1. init: setup
// 2. get Validators
// 3. calculation ValidatorExa and delayedExitTokenIds
// 4. clVaultBalance
// 5. calculationExitValidatorInfo
// 6. calculationForOperator for rewards
// 7. dealLargeExitDelayedRequest for LargeExitDelayedRequestIds
// 8. obtainReportData
func (v *WithdrawHelper) buildReportData(ctx context.Context) error {
	if err := v.setup(ctx); err != nil {
		return errors.Wrap(err, "")
	}

	if err := v.obtainValidatorConsensusInfo(ctx); err != nil {
		return errors.Wrap(err, "")
	}

	if err := v.calculationValidatorExa(ctx); err != nil {
		return errors.Wrap(err, "")
	}

	if err := v.calculationClVaultBalance(ctx); err != nil {
		return errors.Wrap(err, "")
	}

	if err := v.calculationExitValidatorInfo(ctx); err != nil {
		return errors.Wrap(err, "")
	}

	if err := v.calculationForOperator(ctx); err != nil {
		return errors.Wrap(err, "")
	}

	if err := v.dealLargeExitDelayedRequest(ctx); err != nil {
		return errors.Wrap(err, "")
	}

	if err := v.obtainReportData(ctx); err != nil {
		return errors.Wrap(err, "")
	}

	return nil
}

func (v *WithdrawHelper) processReportData(ctx context.Context, reportHash [32]byte) error {
	_, memberInfo, err := v.hashConsensusHelper.GetLastData(ctx)
	if err != nil {
		return errors.Wrap(err, "")
	}

	if memberInfo.CurrentFrameConsensusReport == eth1.ZERO_HASH {
		logger.Info("Quorum is not ready.")
		time.Sleep(time.Second * SECONDS_PER_SLOT)
		return nil
	}

	if memberInfo.CurrentFrameConsensusReport != reportHash {
		logger.Error("Oracle`s hash differs from consensus report hash.")
		return errors.New("Oracle`s hash differs from consensus report hash.")
	}

	submitted, err := v.oracle.IsMainDataSubmitted(ctx)
	if err != nil {
		return errors.Wrap(err, "")
	}
	if submitted {
		logger.Info("Main data already submitted.")
		time.Sleep(time.Second * SECONDS_PER_EPOCH)
	}

	logger.Info("Sending report data...")

	tx, err := contracts.WithdrawOracleContract.Contract.SubmitReportData(v.keyTransactOpts, *v.reportData, v.consensusVersion)
	if err != nil {
		return errors.Wrap(err, "WithdrawOracle SubmitReportData err.")
	}
	logger.Info("Send report data success.",
		zap.String("tx hash", tx.Hash().String()),
		zap.String("consensusVersion", v.consensusVersion.String()),
	)

	return nil
}

func (v *WithdrawHelper) setup(ctx context.Context) error {
	var err error

	chainID, err := eth1.ElClient.Client.ChainID(ctx)
	if err != nil {
		return errors.Wrap(err, "Failed to get chainID.")
	}

	v.keyTransactOpts, err = eth1.KeyTransactOpts(chainID)
	if err != nil {
		return errors.Wrap(err, "Failed to get KeyTransactOpts by privateKey.")
	}

	// delayedExitSlashStandard
	v.delayedExitSlashStandard, err = contracts.OperatorSlashContract.Contract.DelayedExitSlashStandard(nil)
	if err != nil {
		return errors.Wrap(err, "Failed to get OperatorSlashContract delayedExitSlashStandard.")
	}

	// clVaultMinSettleLimit
	v.clVaultMinSettleLimit, err = contracts.WithdrawOracleContract.Contract.ClVaultMinSettleLimit(nil)
	if err != nil {
		return errors.Wrap(err, "Failed to get WithdrawOracleContract clVaultMinSettleLimit.")
	}

	// exitRequestLimit
	v.exitRequestLimit, err = contracts.WithdrawOracleContract.Contract.ExitRequestLimit(nil)
	if err != nil {
		return errors.Wrap(err, "Failed to get WithdrawOracleContract ExitRequestLimit.")
	}

	v.totalNftCount, err = contracts.VnftContract.Contract.TotalSupply(nil)
	if err != nil {
		return errors.Wrap(err, "Failed to get VnftContract TotalSupply.")
	}

	v.delayedExitTokenIds = make([]*big.Int, 0)
	v.largeExitDelayedRequestIds = make([]*big.Int, 0)
	v.withdrawInfos = make([]withdrawOracle.WithdrawInfo, 0)
	v.exitValidatorInfos = make([]withdrawOracle.ExitValidatorInfo, 0)

	v.oracle = &Oracle{}
	v.hashConsensusHelper = &consensusModule.HashConsensusHelper{
		ReportContract:  v.oracle,
		KeyTransactOpts: v.keyTransactOpts,
	}

	consensusVersion, err := contracts.WithdrawOracleContract.Contract.GetConsensusVersion(nil)
	if err != nil {
		return errors.Wrap(err, "Failed to get WithdrawOracleContract GetConsensusVersion.")
	}
	v.consensusVersion = consensusVersion

	refSlot, canReport, err := v.hashConsensusHelper.GetRefSlotAndIsReport(ctx)
	v.refSlot = refSlot

	// executionBlock
	executionBlock, err := consensus.ConsensusClient.CustomizeBeaconService.ExecutionBlock(ctx, v.refSlot.String())
	if err != nil {
		return errors.Wrapf(err, "Failed to get execution block. slot:%s", v.refSlot.String())
	}
	v.executionBlock = executionBlock

	if err != nil {
		return errors.Wrap(err, "")
	}
	if !canReport {
		time.Sleep(time.Second * SECONDS_PER_EPOCH)
	}

	//  init
	v.clBalance = big.NewInt(0)
	v.clSettleAmount = big.NewInt(0)
	v.totalOperatorClCapital = big.NewInt(0)
	v.totalNftCount = big.NewInt(0)

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
		ConsensusVersion:           v.consensusVersion,
		RefSlot:                    v.refSlot,
		ClBalance:                  v.clBalance,
		ClVaultBalance:             v.clVaultBalance,
		ReportExitedCount:          big.NewInt(int64(len(v.exitValidatorInfos))),
		WithdrawInfos:              v.withdrawInfos,
		ExitValidatorInfos:         v.exitValidatorInfos,
		DelayedExitTokenIds:        v.delayedExitTokenIds,
		LargeExitDelayedRequestIds: v.largeExitDelayedRequestIds,
		ClSettleAmount:             v.clSettleAmount,
	}

	return nil
}
