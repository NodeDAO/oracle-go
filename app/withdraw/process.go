// desc:
// @author renshiwei
// Date: 2023/4/10 15:03

package withdraw

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/NodeDAO/oracle-go/app/consensusModule"
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

	if err := v.buildReportData(ctx); err != nil {
		return errors.Wrap(err, "buildReportData err.")
	}

	withdrawOracleReportDataHash, err := EncodeReportData(v.reportData)
	if err != nil {
		return errors.Wrap(err, "EncodeReportData err.")
	}

	reportHashArr := make([][32]byte, 0)
	reportHashArr = append(reportHashArr, withdrawOracleReportDataHash)
	// todo
	reportHashArr = append(reportHashArr, eth1.ZERO_HASH)

	if err := v.hashConsensusHelper.ProcessReportHash(ctx, reportHashArr, v.refSlot, v.reportData); err != nil {
		return errors.Wrap(err, "ProcessReportHash err.")
	}

	if err := v.processReportData(ctx, reportHashArr); err != nil {
		return errors.Wrap(err, "processReportData err.")
	}

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
		return errors.Wrap(err, "buildReportData err.")
	}

	if err := v.obtainValidatorConsensusInfo(ctx); err != nil {
		return errors.Wrap(err, "buildReportData err.")
	}

	if err := v.calculationValidatorExa(ctx); err != nil {
		return errors.Wrap(err, "buildReportData err.")
	}

	if err := v.calculationClVaultBalance(ctx); err != nil {
		return errors.Wrap(err, "buildReportData err.")
	}

	if err := v.calculationExitValidatorInfo(ctx); err != nil {
		return errors.Wrap(err, "buildReportData err.")
	}

	if err := v.calculationForOperator(ctx); err != nil {
		return errors.Wrap(err, "buildReportData err.")
	}

	if err := v.dealLargeExitDelayedRequest(ctx); err != nil {
		return errors.Wrap(err, "buildReportData err.")
	}

	if err := v.obtainReportData(ctx); err != nil {
		return errors.Wrap(err, "buildReportData err.")
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

	submitted, err := v.oracle.IsMainDataSubmitted(ctx)
	if err != nil {
		return errors.Wrap(err, "")
	}
	if submitted {
		logger.Debug("Main data already submitted.")
		return errs.NewSleepError("Main data already submitted.", RandomSleepTime())
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
		// todo
		err := v.oracle.simulatedSubmitReportData(ctx, v.keyTransactOpts, *v.reportData, v.consensusVersion)
		if err != nil {
			return errors.Wrap(err, "simulatedSubmitReportData err.")
		}

		return errs.NewSleepError("simulated report data success.", RandomSleepTime())
	}

	logger.Debug("Sending report data...")

	//opt := v.keyTransactOpts
	//opt.GasLimit = 2000000
	// todo v.consensusVersion => contractVersion
	tx, err := contracts.WithdrawOracleContract.Contract.SubmitReportData(v.keyTransactOpts, *v.reportData, v.consensusVersion, v.withdrawOracleModuleId)
	if err != nil {
		return errors.Wrap(err, "WithdrawOracle SubmitReportData err.")
	}
	// Wait for the transaction to complete
	if _, err = bind.WaitMined(ctx, eth1.ElClient.Client, tx); err != nil {
		return errors.Wrapf(err, "Failed to WaitMined submit report data. tx hash:%s", tx.Hash().String())
	}
	logger.Info("Send report data success.",
		zap.String("refSlot", v.refSlot.String()),
		zap.String("tx hash", tx.Hash().String()),
		zap.String("from", v.keyTransactOpts.From.String()),
		zap.String("to", contracts.WithdrawOracleContract.Address),
		zap.Uint64("gas", tx.Gas()),
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

	v.withdrawOracleModuleId, err = v.hashConsensusHelper.GetModuleId(ctx, common.HexToAddress(contracts.WithdrawOracleContract.Address))
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
