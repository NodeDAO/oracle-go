// desc:
// @author renshiwei
// Date: 2023/4/10 15:03

package withdraw

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/NodeDAO/oracle-go/app/consensusModule"
	"github.com/NodeDAO/oracle-go/app/largestake"
	"github.com/NodeDAO/oracle-go/common/errs"
	"github.com/NodeDAO/oracle-go/common/logger"
	"github.com/NodeDAO/oracle-go/config"
	"github.com/NodeDAO/oracle-go/consensus"
	"github.com/NodeDAO/oracle-go/contracts"
	"github.com/NodeDAO/oracle-go/contracts/withdrawOracle"
	"github.com/NodeDAO/oracle-go/eth1"
	"github.com/NodeDAO/oracle-go/utils/typetool"
	consensusclient "github.com/attestantio/go-eth2-client"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
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
		return errs.NewSleepError("withdrawOracle is paused.", RandomSleepTime())
	}

	if err := v.check(ctx); err != nil {
		return errors.Wrap(err, "check network、el、cl config err.")
	}

	if err := v.setup(ctx); err != nil {
		return errors.Wrap(err, "setup err.")
	}

	if err := v.reportHashArr(ctx); err != nil {
		return errors.Wrap(err, "reportHashArr err.")
	}

	if err := v.hashConsensusHelper.ProcessReportHash(ctx, v.consensusReportHashArr, v.refSlot, v.reportData); err != nil {
		return errors.Wrap(err, "ProcessReportHash err.")
	}

	if err := v.processReportData(ctx, v.consensusReportHashArr); err != nil {
		return errors.Wrap(err, "processReportData err.")
	}

	return nil
}

func (v *WithdrawHelper) reportHashArr(ctx context.Context) error {
	var err error
	v.isWithdrawOracleNeedReport, err = v.hashConsensusHelper.IsModuleReport(big.NewInt(WITHRAW_ORACLE_MODULE_ID), v.refSlot)
	if err != nil {
		return errors.Wrap(err, "[WithdrawOracle] IsModuleReport err.")
	}
	v.isLargeStakeOracleNeedReport, err = v.hashConsensusHelper.IsModuleReport(big.NewInt(LARGE_STAKE_MODULE_ID), v.refSlot)
	if err != nil {
		return errors.Wrap(err, "[LargeStakeOracle] IsModuleReport err.")
	}

	if v.isWithdrawOracleNeedReport {
		if err := v.buildWithdrawOracleReportData(ctx); err != nil {
			return errors.Wrap(err, "[WithdrawOracle] buildWithdrawOracleReportData err.")
		}

		withdrawOracleReportDataHash, err := EncodeReportData(v.reportData)
		if err != nil {
			return errors.Wrap(err, "[WithdrawOracle] EncodeReportData err.")
		}
		v.consensusReportHashArr = append(v.consensusReportHashArr, withdrawOracleReportDataHash)
	} else {
		logger.Debug("[WithdrawOracle] not need report.")
		v.consensusReportHashArr = append(v.consensusReportHashArr, eth1.ZERO_HASH)
	}

	if v.isLargeStakeOracleNeedReport {
		v.largeStakeOracleRes, err = v.largeStakeOracleHelper.ProcessReportData(ctx)
		if err != nil {
			return errors.Wrap(err, "[LargeStakeOracle] ProcessReportData err.")
		}

		v.consensusReportHashArr = append(v.consensusReportHashArr, v.largeStakeOracleRes.ReportDataHash)
		if !v.largeStakeOracleRes.IsNeedReport {
			v.isLargeStakeOracleNeedReport = false
		}
	} else {
		logger.Debug("[LargeStakeOracle] not need report.")
		v.consensusReportHashArr = append(v.consensusReportHashArr, eth1.ZERO_HASH)
	}

	if !v.isWithdrawOracleNeedReport && !v.isLargeStakeOracleNeedReport {
		v.isConsensusReport = false
		return errs.NewSleepError("All Oracle not need to report consensus.", RandomSleepTime())
	} else {
		v.isConsensusReport = true
	}

	return nil
}

// buildWithdrawOracleReportData core process
// 1. get Validators
// 2. calculation ValidatorExa
// 3. clVaultBalance
// 4. calculationExitValidatorInfo
// 5. calculationForOperator for rewards
// 6. obtainWithdrawOracleReportData
func (v *WithdrawHelper) buildWithdrawOracleReportData(ctx context.Context) error {

	if err := v.obtainValidatorConsensusInfo(ctx); err != nil {
		return errors.Wrap(err, "buildWithdrawOracleReportData err.")
	}

	if err := v.calculationValidatorExa(ctx); err != nil {
		return errors.Wrap(err, "buildWithdrawOracleReportData err.")
	}

	if err := v.calculationClVaultBalance(ctx); err != nil {
		return errors.Wrap(err, "buildWithdrawOracleReportData err.")
	}

	if err := v.calculationExitValidatorInfo(ctx); err != nil {
		return errors.Wrap(err, "buildWithdrawOracleReportData err.")
	}

	if err := v.calculationForOperator(ctx); err != nil {
		return errors.Wrap(err, "buildWithdrawOracleReportData err.")
	}

	if err := v.obtainWithdrawOracleReportData(ctx); err != nil {
		return errors.Wrap(err, "buildWithdrawOracleReportData err.")
	}

	return nil
}

func (v *WithdrawHelper) processReportData(ctx context.Context, reportHash [][32]byte) error {
	// If the configuration does not report real data, return it directly
	if !config.Config.Oracle.IsReportData {
		return nil
	}

	_, memberInfo, err := v.hashConsensusHelper.GetLastData(ctx)
	if err != nil {
		return errors.Wrap(err, "hashConsensus GetLastData err.")
	}

	if len(memberInfo.CurrentFrameConsensusReport) == 0 {
		logger.Debug("Quorum is not ready.")
		return errs.NewSleepError("Quorum is not ready.", RandomSleepTime())
	}

	if !typetool.CompareByte32Arrays(memberInfo.CurrentFrameConsensusReport, reportHash) {
		logger.Error("Oracle`s hash differs from consensus report hash.")
		return errors.New("Oracle`s hash differs from consensus report hash.")
	}

	submitted, err := v.checkALlOracleDataSubmitted()
	if err != nil {
		return errors.Wrap(err, "checkALlOracleDataSubmitted err.")
	}
	if submitted {
		return errs.NewSleepError("All Oracle main data already submitted.", RandomSleepTime())
	}

	reportJson, err := json.Marshal(v.reportData)
	if err != nil {
		logger.Debug("report data.", zap.String("report data", fmt.Sprintf("%+v", v.reportData)))
	} else {
		logger.Debug("report data.", zap.String("report data", fmt.Sprintf("%s", string(reportJson))))
	}

	// If configured to only simulate transactions
	if config.Config.Oracle.IsSimulatedReportData {
		logger.Debug("simulated report data...")
		if v.isWithdrawOracleNeedReport {
			err := v.oracle.simulatedWithdrawOracleSubmitReportData(ctx, v.keyTransactOpts, *v.reportData, big.NewInt(WITHRAW_ORACLE_CONTRACT_VERSION), big.NewInt(WITHRAW_ORACLE_MODULE_ID))
			if err != nil {
				return errors.Wrap(err, "[WithdrawOracle] simulatedWithdrawOracleSubmitReportData err.")
			}
		}

		if v.isLargeStakeOracleNeedReport {
			err = v.oracle.simulatedLargeStakeOracleSubmitReportData(ctx, v.keyTransactOpts, v.largeStakeOracleRes.ReportData, big.NewInt(LARGE_STAKE_ORACLE_CONTRACT_VERSION), big.NewInt(LARGE_STAKE_MODULE_ID))
			if err != nil {
				return errors.Wrap(err, "[WithdrawOracle] simulatedWithdrawOracleSubmitReportData err.")
			}
		}

		return errs.NewSleepError("simulated report data success.", RandomSleepTime())
	}

	logger.Debug("Sending report data...")

	opt := v.keyTransactOpts
	opt.GasLimit = 2000000
	if v.isWithdrawOracleNeedReport {
		tx, err := contracts.WithdrawOracleContract.Contract.SubmitReportData(opt, *v.reportData, big.NewInt(WITHRAW_ORACLE_CONTRACT_VERSION), big.NewInt(WITHRAW_ORACLE_MODULE_ID))
		if err != nil {
			return errors.Wrap(err, "[WithdrawOracle] SubmitReportData err.")
		}
		// Wait for the transaction to complete
		if _, err = bind.WaitMined(ctx, eth1.ElClient.Client, tx); err != nil {
			return errors.Wrapf(err, "[WithdrawOracle] Failed to WaitMined submit report data. tx hash:%s", tx.Hash().String())
		}
		logger.Info("[WithdrawOracle] Send report data success.",
			zap.String("refSlot", v.refSlot.String()),
			zap.String("tx hash", tx.Hash().String()),
			zap.String("from", v.keyTransactOpts.From.String()),
			zap.String("to", contracts.WithdrawOracleContract.Address),
			zap.Uint64("gas", tx.Gas()),
			zap.String("consensusVersion", v.consensusVersion.String()),
		)
	}

	if v.isLargeStakeOracleNeedReport {
		tx, err := contracts.LargeStakeOracleContract.Contract.SubmitReportData(v.keyTransactOpts, v.largeStakeOracleRes.ReportData, big.NewInt(LARGE_STAKE_ORACLE_CONTRACT_VERSION), big.NewInt(LARGE_STAKE_MODULE_ID))
		if err != nil {
			return errors.Wrap(err, "[LargeStakeOracle] SubmitReportData err.")
		}
		// Wait for the transaction to complete
		if _, err = bind.WaitMined(ctx, eth1.ElClient.Client, tx); err != nil {
			return errors.Wrapf(err, "[LargeStakeOracle] Failed to WaitMined submit report data. tx hash:%s", tx.Hash().String())
		}
		logger.Info("[LargeStakeOracle] Send report data success.",
			zap.String("refSlot", v.refSlot.String()),
			zap.String("tx hash", tx.Hash().String()),
			zap.String("from", v.keyTransactOpts.From.String()),
			zap.String("to", contracts.WithdrawOracleContract.Address),
			zap.Uint64("gas", tx.Gas()),
			zap.String("consensusVersion", v.consensusVersion.String()),
		)
	}

	return nil
}

func (v *WithdrawHelper) checkALlOracleDataSubmitted() (bool, error) {
	withdrawOracleProcessingState, err := contracts.WithdrawOracleContract.Contract.GetProcessingState(nil)
	if err != nil {
		return false, errors.Wrap(err, "withdrawOracleContract GetProcessingState err.")
	}
	largeStakeOracleProcessingState, err := contracts.LargeStakeOracleContract.Contract.GetProcessingState(nil)
	if err != nil {
		return false, errors.Wrap(err, "LargeStakeOracleContract GetProcessingState err.")
	}

	if withdrawOracleProcessingState.DataSubmitted && largeStakeOracleProcessingState.DataSubmitted {
		logger.Debug("[ALL Oracle] Main data already submitted.")
		return true, nil
	}
	if withdrawOracleProcessingState.DataSubmitted && !v.isLargeStakeOracleNeedReport {
		logger.Debug("[WithdrawOracle] main data already submitted. [LargeStakeOracle] not need to report.")
		return true, nil
	}
	if largeStakeOracleProcessingState.DataSubmitted && !v.isWithdrawOracleNeedReport {
		logger.Debug("[LargeStakeOracle] main data already submitted. [WithdrawOracle] not need to report.")
		return true, nil
	}

	return false, nil
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

	v.withdrawInfos = make([]withdrawOracle.WithdrawInfo, 0)
	v.exitValidatorInfos = make([]withdrawOracle.ExitValidatorInfo, 0)
	v.consensusReportHashArr = make([][32]byte, 0)

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

	v.withdrawOracleModuleId, err = v.hashConsensusHelper.GetModuleId(common.HexToAddress(contracts.WithdrawOracleContract.Address))
	if err != nil {
		return errors.Wrap(err, "Failed to withdrawOracleModuleId.")
	}

	refSlot, deadlineSlot, err := v.hashConsensusHelper.GetRefSlotAndIsReport(ctx)
	if err != nil {
		return errors.Wrap(err, "Failed to GetRefSlotAndIsReport. slot:head")
	}

	v.refSlot = refSlot
	v.deadlineSlot = deadlineSlot

	// executionBlock
	executionBlock, err := consensus.ConsensusClient.CustomizeBeaconService.ExecutionBlock(ctx, v.refSlot.String())
	if err != nil {
		return errors.Wrapf(err, "Failed to get execution block. slot:%s", v.refSlot.String())
	}
	v.executionBlock = executionBlock
	logger.Debug("execution block",
		zap.String("BlockNumber", executionBlock.BlockNumber.String()),
		zap.String("BlockHash", executionBlock.BlockHash),
		zap.String("Timestamp", executionBlock.Timestamp),
	)

	if err != nil {
		return errors.Wrap(err, "setup err.")
	}

	v.largeStakeOracleHelper = largestake.NewLargeStakeHelper(v.refSlot, big.NewInt(CONSENSUS_VERSION))

	return nil
}

func (v *WithdrawHelper) check(ctx context.Context) error {
	// check network
	chainID, err := eth1.ElClient.Client.ChainID(ctx)
	isSameNetwork := eth1.IsSameNetwork(config.Config.Eth.Network, int(chainID.Int64()))
	if !isSameNetwork {
		return errs.NewSleepError("el is not same network.", RandomSleepTime())
	}

	depositContract, err := consensus.ConsensusClient.ConsensusClient.(consensusclient.DepositContractProvider).DepositContract(ctx)
	if err != nil {
		return errors.Wrap(err, "get DepositContract err.")
	}
	isSameNetwork = eth1.IsSameNetwork(config.Config.Eth.Network, int(depositContract.ChainID))
	if !isSameNetwork {
		return errs.NewSleepError("cl is not same network.", RandomSleepTime())
	}

	// check beacon sync
	syncState, err := consensus.ConsensusClient.ConsensusClient.(consensusclient.NodeSyncingProvider).NodeSyncing(ctx)
	if err != nil {
		return errors.Wrap(err, "check err.")
	}
	if syncState.IsSyncing {
		logger.Warnf("Beacon node is syncing. isSync:%v SyncDistance:%v", syncState.IsSyncing, syncState.SyncDistance)
		return errs.NewSleepError("Beacon node is syncing.", RandomSleepTime())
	}

	// check el sync
	blockHeader, err := consensus.ConsensusClient.ConsensusClient.(consensusclient.BeaconBlockHeadersProvider).BeaconBlockHeader(ctx, "finalized")
	if err != nil {
		return errors.Wrap(err, "get BeaconBlockHeader finalized err.")
	}
	finalizedSlot := blockHeader.Header.Message.Slot

	consensusBlock, err := consensus.ConsensusClient.CustomizeBeaconService.ExecutionBlock(ctx, fmt.Sprintf("%v", finalizedSlot))
	if err != nil {
		return errors.Wrap(err, "get ExecutionBlock err.")
	}

	blockNumber, err := eth1.ElClient.Client.BlockNumber(ctx)

	if blockNumber < consensusBlock.BlockNumber.Uint64() {
		logger.Warn("El node is syncing.",
			zap.Uint64("beacon blockNumber", consensusBlock.BlockNumber.Uint64()),
			zap.Uint64("el blockNumber", blockNumber),
		)
		return errs.NewSleepError("el node blockNumber is old.", RandomSleepTime())
	}

	if blockNumber > consensusBlock.BlockNumber.Uint64()+100 {
		logger.Warn("el and cl is not same env.",
			zap.Uint64("beacon blockNumber", consensusBlock.BlockNumber.Uint64()),
			zap.Uint64("el blockNumber", blockNumber),
		)
		return errs.NewSleepError("el and cl is not same env.", RandomSleepTime())
	}

	return nil
}

// Slice types need to be sorted
func (v *WithdrawHelper) obtainWithdrawOracleReportData(ctx context.Context) error {
	// sort slice

	sort.Slice(v.withdrawInfos, func(i, j int) bool {
		return v.withdrawInfos[i].OperatorId < v.withdrawInfos[j].OperatorId
	})

	sort.Slice(v.exitValidatorInfos, func(i, j int) bool {
		return v.exitValidatorInfos[i].ExitTokenId < v.exitValidatorInfos[j].ExitTokenId
	})

	reportExitedCount := big.NewInt(0)
	if len(v.exitValidatorInfos) > 0 {
		reportExitedCount = big.NewInt(int64(len(v.exitValidatorInfos)))
	}

	v.reportData = &withdrawOracle.WithdrawOracleReportData{
		ConsensusVersion:   v.consensusVersion,
		RefSlot:            v.refSlot,
		ClBalance:          v.clBalance,
		ClVaultBalance:     v.clVaultBalance,
		ReportExitedCount:  reportExitedCount,
		WithdrawInfos:      v.withdrawInfos,
		ExitValidatorInfos: v.exitValidatorInfos,
		ClSettleAmount:     v.clSettleAmount,
	}

	return nil
}
