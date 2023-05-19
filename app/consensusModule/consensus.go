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
	"github.com/NodeDAO/oracle-go/contracts/withdrawOracle"
	"github.com/NodeDAO/oracle-go/eth1"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"math/big"
)

func (v *HashConsensusHelper) ProcessReportHash(ctx context.Context, dataHash [32]byte, refSlot, consensusVersion *big.Int, reportData *withdrawOracle.WithdrawOracleReportData) error {
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
	isConsensusReportAlreadyProcessing, err := v.isConsensusReportAlreadyProcessing(ctx, memberInfo)
	if err != nil {
		return errors.Wrap(err, "get isConsensusReportAlreadyProcessing err.")
	}
	if isConsensusReportAlreadyProcessing {
		return errs.NewSleepError("Consensus Report Already Processing (Main data already submitted)", RandomSleepTime())
	}

	dataHashStr := hexutil.Encode(dataHash[:])
	if dataHash != memberInfo.CurrentFrameMemberReport {
		oldDataHashStr := hexutil.Encode(memberInfo.CurrentFrameMemberReport[:])
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

		err := v.submitReport(ctx, dataHash, refSlot, consensusVersion)
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

func (v *HashConsensusHelper) submitReport(ctx context.Context, dataHash [32]byte, refSlot, consensusVersion *big.Int) error {
	consensusContract, err := v.ReportContract.GetConsensusContract(ctx)
	if err != nil {
		return errors.Wrap(err, "")
	}

	//opt := v.KeyTransactOpts
	//opt.GasLimit = 2000000
	tx, err := consensusContract.SubmitReport(v.KeyTransactOpts, refSlot, dataHash, consensusVersion)
	if err != nil {
		return errors.Wrapf(err, "Failed to submit consensus report. refslot:%s, consensusVersion:%s", refSlot.String(), consensusVersion.String())
	}
	// Wait for the transaction to complete
	if _, err = bind.WaitMined(context.Background(), eth1.ElClient.Client, tx); err != nil {
		return errors.Wrapf(err, "Failed to WaitMined submit consensus report. refslot:%s, consensusVersion:%s", refSlot.String(), consensusVersion.String())
	}

	to, err := v.ReportContract.GetConsensusContractAddress(ctx)
	if err != nil {
		processorAddress, _ := v.ReportContract.GetReportAsyncProcessorAddress()
		logger.Warn("GetConsensusContractAddress failed.", zap.String("ReportAsyncProcessor", processorAddress.String()))
	}

	logger.Info("submit consensus Report success.",
		zap.String("refSlot", refSlot.String()),
		zap.String("tx hash", tx.Hash().String()),
		zap.String("from", v.KeyTransactOpts.From.String()),
		zap.String("to", to.String()),
		zap.Uint64("gas", tx.Gas()),
		zap.String("consensusVersion", consensusVersion.String()),
	)

	return nil
}

func (v *HashConsensusHelper) GetMemberInfo(ctx context.Context) (*MemberInfo, error) {
	consensusContract, err := v.ReportContract.GetConsensusContract(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	currentFrame, err := consensusContract.GetCurrentFrame(nil)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to get HashConsensus's GetCurrentFrame.")
	}

	frameConfig, err := consensusContract.GetFrameConfig(nil)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to get HashConsensus's GetCurrentFrame.")
	}

	// v.KeyTransactOpts.From
	// mainnet address
	//mainnetAddress := common.HexToAddress("0x080C185D164446746068Db1650850F453ffdB92c")
	memberConsensusState, err := consensusContract.GetConsensusStateForMember(nil, v.KeyTransactOpts.From)
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

func (v *HashConsensusHelper) isConsensusReportAlreadyProcessing(ctx context.Context, memberInfo *MemberInfo) (bool, error) {
	lastProcessingRefSlot, err := v.ReportContract.GetLastProcessingRefSlot(ctx)
	if err != nil {
		return false, errors.Wrap(err, "")
	}

	if memberInfo.CurrentFrameRefSlot.Cmp(lastProcessingRefSlot) != 1 {
		if memberInfo.CurrentFrameRefSlot.Cmp(memberInfo.LastReportRefSlot) == 0 {
			return true, nil
		}
	}

	return false, nil
}
