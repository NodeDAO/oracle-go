// desc:
// @author renshiwei
// Date: 2023/4/10 15:03

package withdraw

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/NodeDAO/oracle-go/app/consensusModule"
	"github.com/NodeDAO/oracle-go/common/logger"
	"github.com/NodeDAO/oracle-go/config"
	"github.com/NodeDAO/oracle-go/consensus"
	"github.com/NodeDAO/oracle-go/contracts"
	"github.com/NodeDAO/oracle-go/contracts/withdrawOracle"
	"github.com/NodeDAO/oracle-go/eth1"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"math/big"
	"sort"
)

func (v *WithdrawHelper) ProcessReport(ctx context.Context) error {
	paused, err := v.oracle.Paused(ctx)
	if err != nil {
		return errors.Wrap(err, "")
	}
	if paused {
		logger.Info("withdrawOracle is paused.")
		DefaultRandomSleep()
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
	// If the configuration does not report real data, return it directly
	if !config.Config.Oracle.IsReportData {
		return nil
	}

	_, memberInfo, err := v.hashConsensusHelper.GetLastData(ctx)
	if err != nil {
		return errors.Wrap(err, "")
	}

	if memberInfo.CurrentFrameConsensusReport == eth1.ZERO_HASH {
		logger.Info("Quorum is not ready.")
		DefaultRandomSleep()
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
		DefaultRandomSleep()
	}

	logger.Info("Sending report data...")

	reportJson, err := json.Marshal(v.reportData)
	if err != nil {
		logger.Debug("report data.", zap.String("report data", fmt.Sprintf("%+v", v.reportData)))
	} else {
		logger.Debug("report data.", zap.String("report data", fmt.Sprintf("%s", string(reportJson))))
	}

	// If configured to only simulate transactions
	if config.Config.Oracle.IsSimulatedReportData {
		err := v.oracle.simulatedSubmitReportData(ctx, *v.reportData, v.consensusVersion)
		if err != nil {
			return errors.Wrap(err, "")
		}

		return nil
	}

	//opt := v.keyTransactOpts
	//opt.GasLimit = 2000000
	tx, err := contracts.WithdrawOracleContract.Contract.SubmitReportData(v.keyTransactOpts, *v.reportData, v.consensusVersion)
	if err != nil {
		return errors.Wrap(err, "WithdrawOracle SubmitReportData err.")
	}
	// Wait for the transaction to complete
	if _, err = bind.WaitMined(context.Background(), eth1.ElClient.Client, tx); err != nil {
		return errors.Wrapf(err, "Failed to WaitMined submit report data. tx hash:%s", tx.Hash().String())
	}
	logger.Info("Send report data success.",
		zap.String("tx hash", tx.Hash().String()),
		zap.String("consensusVersion", v.consensusVersion.String()),
	)

	return nil
}

func (v *WithdrawHelper) setup(ctx context.Context) error {
	//  init
	v.clBalance = big.NewInt(0)
	v.clSettleAmount = big.NewInt(0)
	v.totalOperatorClCapital = big.NewInt(0)
	v.totalNftCount = big.NewInt(0)

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
	if err != nil {
		return errors.Wrap(err, "Failed to GetRefSlotAndIsReport. slot:head")
	}

	v.refSlot = refSlot
	logger.Debug("Oracle start scan ...", zap.String("refSlot", refSlot.String()))

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
		DefaultRandomSleep()
	}

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

	reportExitedCount := big.NewInt(0)
	if len(v.exitValidatorInfos) > 0 {
		reportExitedCount = big.NewInt(int64(len(v.exitValidatorInfos)))
	}

	v.reportData = &withdrawOracle.WithdrawOracleReportData{
		ConsensusVersion:           v.consensusVersion,
		RefSlot:                    v.refSlot,
		ClBalance:                  v.clBalance,
		ClVaultBalance:             v.clVaultBalance,
		ReportExitedCount:          reportExitedCount,
		WithdrawInfos:              v.withdrawInfos,
		ExitValidatorInfos:         v.exitValidatorInfos,
		DelayedExitTokenIds:        v.delayedExitTokenIds,
		LargeExitDelayedRequestIds: v.largeExitDelayedRequestIds,
		ClSettleAmount:             v.clSettleAmount,
	}

	return nil
}
