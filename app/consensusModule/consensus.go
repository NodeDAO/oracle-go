// desc:
// @author renshiwei
// Date: 2023/4/11 15:28

package consensusModule

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/NodeDAO/oracle-go/common/errs"
	"github.com/NodeDAO/oracle-go/common/logger"
	"github.com/NodeDAO/oracle-go/config"
	"github.com/NodeDAO/oracle-go/consensus"
	"github.com/NodeDAO/oracle-go/contracts"
	"github.com/NodeDAO/oracle-go/contracts/withdrawOracle"
	"github.com/NodeDAO/oracle-go/eth1"
	"github.com/NodeDAO/oracle-go/utils/typetool"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"math/big"
	"strings"
)

func (v *HashConsensusHelper) ProcessReportHash(ctx context.Context, dataHash [][32]byte, refSlot *big.Int, reportData *withdrawOracle.WithdrawOracleReportData) error {
	headSlot, err := consensus.ConsensusClient.CustomizeBeaconService.HeadSlot(ctx)
	if err != nil {
		return errors.Wrap(err, "ProcessReportHash HeadSlot err.")
	}

	memberInfo, err := v.GetMemberInfo(ctx)
	if err != nil {
		return errors.Wrap(err, "ProcessReportHash GetMemberInfo err.")
	}
	if !memberInfo.IsFastLane {
		if headSlot.Cmp(new(big.Int).Add(memberInfo.CurrentFrameRefSlot, memberInfo.FastLaneLengthSlot)) == -1 {
			msg := fmt.Sprintf("Member is not in fast lane, so report will be postponed for %s slots", memberInfo.FastLaneLengthSlot.String())
			return errs.NewSleepError(msg, RandomSleepTime())
		}
	}

	if !memberInfo.IsReportMember {
		msg := fmt.Sprintf("Account can`t submit report hash.Not Oracle member. Account:%s", v.KeyTransactOpts.From.String())
		return errs.NewSleepError(msg, RandomSleepTime())
	}

	// If Oracle has already submitted real data, it is not submitting consensus
	isConsensusReportAlreadyProcessing, err := v.isConsensusReportAlreadyProcessing()
	if err != nil {
		return errors.Wrap(err, "get isConsensusReportAlreadyProcessing err.")
	}
	if isConsensusReportAlreadyProcessing {
		return errs.NewSleepError("Consensus Report Already Processing (Main data already submitted)", RandomSleepTime())
	}

	dataHashStr := strings.Join(typetool.Byte32ArrToStrArr(dataHash), ",")

	if !typetool.CompareByte32Arrays(dataHash, memberInfo.CurrentFrameMemberReport) {
		oldDataHashStr := strings.Join(typetool.Byte32ArrToStrArr(memberInfo.CurrentFrameMemberReport), ",")
		reportJson, _ := json.Marshal(reportData)

		if memberInfo.IsCurrentReportConsensus && !config.Config.Oracle.IsDifferentConsensusHashReport {
			logger.Warn("Consensus hash is different.",
				zap.String("new hash", dataHashStr),
				zap.String("old hash", oldDataHashStr),
				zap.String("ReportData", string(reportJson)),
				zap.Bool("IsDifferentConsensusHashReport", config.Config.Oracle.IsDifferentConsensusHashReport),
			)
			return errs.NewSleepError("CurrentFrameMemberReport Consensus hash is different. Member has report consensus.", RandomSleepTime())
		}

		err := v.submitReport(ctx, dataHash, refSlot)
		if err != nil {
			return errors.Wrap(err, "")
		}

		if oldDataHashStr == eth1.ZERO_HASH_STR {
			logger.Debug("Send reportData's hash.", zap.String("hash", dataHashStr), zap.String("ReportData", string(reportJson)))
		} else {
			if config.Config.Oracle.IsDifferentConsensusHashReport {
				logger.Warn("Send reportData's hash. CurrentFrameMemberReport Consensus hash is different.",
					zap.String("new hash", dataHashStr),
					zap.String("old hash", oldDataHashStr),
					zap.String("ReportData", string(reportJson)),
					zap.Bool("IsDifferentConsensusHashReport", config.Config.Oracle.IsDifferentConsensusHashReport),
				)
			}
		}
	} else {
		logger.Debug("Provided hash already submitted.", zap.String("hash", dataHashStr))
	}

	return nil
}

func (v *HashConsensusHelper) submitReport(ctx context.Context, dataHash [][32]byte, refSlot *big.Int) error {
	//opt := v.KeyTransactOpts
	//opt.GasLimit = 2000000
	tx, err := contracts.HashConsensusContract.Contract.SubmitReport(v.KeyTransactOpts, refSlot, dataHash)
	if err != nil {
		return errors.Wrapf(err, "Failed to submit consensus report. refslot:%s", refSlot.String())
	}
	// Wait for the transaction to complete
	if _, err = bind.WaitMined(context.Background(), eth1.ElClient.Client, tx); err != nil {
		return errors.Wrapf(err, "Failed to WaitMined submit consensus report. refslot:%s", refSlot.String())
	}

	logger.Info("submit consensus Report success.",
		zap.String("refSlot", refSlot.String()),
		zap.String("tx hash", tx.Hash().String()),
		zap.String("from", v.KeyTransactOpts.From.String()),
		zap.String("to", tx.To().String()),
		zap.Uint64("gas", tx.Gas()),
	)

	return nil
}

func (v *HashConsensusHelper) GetMemberInfo(ctx context.Context) (*MemberInfo, error) {
	currentFrame, err := contracts.HashConsensusContract.Contract.GetCurrentFrame(nil)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to get HashConsensus's GetCurrentFrame.")
	}

	frameConfig, err := contracts.HashConsensusContract.Contract.GetFrameConfig(nil)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to get HashConsensus's GetCurrentFrame.")
	}

	// v.KeyTransactOpts.From
	// mainnet address
	//mainnetAddress := common.HexToAddress("0x080C185D164446746068Db1650850F453ffdB92c")
	memberConsensusState, err := contracts.HashConsensusContract.Contract.GetConsensusStateForMember(nil, v.KeyTransactOpts.From)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to get HashConsensus's GetConsensusStateForMember.")
	}

	isCurrentReportConsensus := false
	if memberConsensusState.CurrentFrameRefSlot.Cmp(memberConsensusState.LastMemberReportRefSlot) == 0 {
		isCurrentReportConsensus = true
	}

	return &MemberInfo{
		IsReportMember:              memberConsensusState.IsMember,
		IsFastLane:                  memberConsensusState.IsFastLane,
		CanReport:                   memberConsensusState.CanReport,
		IsCurrentReportConsensus:    isCurrentReportConsensus,
		LastReportRefSlot:           memberConsensusState.LastMemberReportRefSlot,
		FastLaneLengthSlot:          frameConfig.FastLaneLengthSlots,
		CurrentFrameRefSlot:         currentFrame.RefSlot,
		DeadlineSlot:                currentFrame.ReportProcessingDeadlineSlot,
		CurrentFrameMemberReport:    memberConsensusState.CurrentFrameMemberReport,
		CurrentFrameConsensusReport: memberConsensusState.CurrentFrameConsensusReport,
	}, nil
}

func (v *HashConsensusHelper) GetRefSlotAndIsReport(ctx context.Context) (*big.Int, *big.Int, error) {
	err := v.ReportContract.CheckContractVersions(ctx)
	if err != nil {
		return nil, nil, errors.Wrap(err, "")
	}

	headSlot, err := consensus.ConsensusClient.CustomizeBeaconService.HeadSlot(ctx)
	if err != nil {
		return nil, nil, errors.Wrap(err, "")
	}

	memberInfo, err := v.GetMemberInfo(ctx)
	if err != nil {
		return nil, nil, errors.Wrap(err, "")
	}

	logger.Debug("withdrawOracle start scan ...", zap.String("refSlot", memberInfo.CurrentFrameRefSlot.String()), zap.String("deadlineSlot", memberInfo.DeadlineSlot.String()))

	if !memberInfo.CanReport {
		return nil, nil, errs.NewSleepError("Member's ConsensusState is not CanReport.", RandomSleepTime())
	}

	// Slot < CurrentFrameRefSlot
	if headSlot.Cmp(memberInfo.CurrentFrameRefSlot) == -1 {
		logger.Warn("Reference slot is not yet finalized.",
			zap.String("headSlot", headSlot.String()),
			zap.String("CurrentFrameRefSlot", memberInfo.CurrentFrameRefSlot.String()))
		return nil, nil, errs.NewSleepError("Reference slot is not yet finalized.", RandomSleepTime())
	}

	// Slot > DeadlineSlot
	if headSlot.Cmp(memberInfo.DeadlineSlot) == 1 {
		logger.Warn("Deadline missed.",
			zap.String("headSlot", headSlot.String()),
			zap.String("DeadlineSlot", memberInfo.DeadlineSlot.String()))
		return nil, nil, errs.NewSleepError("Deadline missed.", RandomSleepTime())
	}

	return memberInfo.CurrentFrameRefSlot, memberInfo.DeadlineSlot, nil
}

func (v *HashConsensusHelper) GetLastData(ctx context.Context) (*big.Int, *MemberInfo, error) {
	headSlot, err := consensus.ConsensusClient.CustomizeBeaconService.HeadSlot(ctx)
	if err != nil {
		return nil, nil, errors.Wrap(err, "")
	}

	memberInfo, err := v.GetMemberInfo(ctx)
	if err != nil {
		return nil, nil, errors.Wrap(err, "")
	}

	return headSlot, memberInfo, nil
}

func (v *HashConsensusHelper) isConsensusReportAlreadyProcessing() (bool, error) {
	withdrawOracleProcessingState, err := contracts.WithdrawOracleContract.Contract.GetProcessingState(nil)
	if err != nil {
		return false, errors.Wrap(err, "withdrawOracleContract GetProcessingState err.")
	}
	largeStakeOracleProcessingState, err := contracts.LargeStakeOracleContract.Contract.GetProcessingState(nil)
	if err != nil {
		return false, errors.Wrap(err, "LargeStakeOracleContract GetProcessingState err.")
	}

	return withdrawOracleProcessingState.DataSubmitted && largeStakeOracleProcessingState.DataSubmitted, nil
}

func (v *HashConsensusHelper) GetModuleId(oracleAddress common.Address) (*big.Int, error) {
	moduleId, err := contracts.HashConsensusContract.Contract.GetReportModuleId(nil, oracleAddress)
	if err != nil {
		return big.NewInt(0), errors.Wrap(err, "consensusContract GetReportModuleId err.")
	}

	if moduleId.Cmp(big.NewInt(0)) == 0 {
		return big.NewInt(0), errors.New("moduleId is zero err.")
	}

	return moduleId, nil
}

func (v *HashConsensusHelper) IsModuleReport(module, slot *big.Int) (bool, error) {
	isNeedReport, err := contracts.HashConsensusContract.Contract.IsModuleReport(nil, module, slot)
	if err != nil {
		return false, errors.Wrap(err, "IsModuleReport err.")
	}
	return isNeedReport, nil
}
