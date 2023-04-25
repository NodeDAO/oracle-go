// desc:
// @author renshiwei
// Date: 2023/4/11 15:28

package consensusModule

import (
	"context"
	"encoding/json"
	"github.com/NodeDAO/oracle-go/common/errs"
	"github.com/NodeDAO/oracle-go/common/logger"
	"github.com/NodeDAO/oracle-go/consensus"
	"github.com/NodeDAO/oracle-go/contracts/withdrawOracle"
	"github.com/NodeDAO/oracle-go/eth1"
	"github.com/NodeDAO/oracle-go/utils/timetool"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"math/big"
	"time"
)

func (v *HashConsensusHelper) ProcessReportHash(ctx context.Context, dataHash [32]byte, refSlot, consensusVersion *big.Int, reportData *withdrawOracle.WithdrawOracleReportData) error {
	headSlot, err := consensus.ConsensusClient.CustomizeBeaconService.HeadSlot(ctx)
	if err != nil {
		return errors.Wrap(err, "")
	}

	memberInfo, err := v.GetMemberInfo(ctx)
	if err != nil {
		return errors.Wrap(err, "")
	}
	if !memberInfo.IsFastLane {
		if headSlot.Cmp(new(big.Int).Add(memberInfo.CurrentFrameRefSlot, memberInfo.FastLaneLengthSlot)) == -1 {
			logger.Infof("Member is not in fast lane, so report will be postponed for %s slots", memberInfo.FastLaneLengthSlot.String())
		}
	}

	if !memberInfo.IsReportMember {
		logger.Warn("Account can`t submit report hash.")
	}

	// If Oracle has already submitted real data, it is not submitting consensus
	isConsensusReportAlreadyProcessing, err := v.isConsensusReportAlreadyProcessing(ctx, memberInfo)
	if err != nil {
		return errors.Wrap(err, "get isConsensusReportAlreadyProcessing err.")
	}
	if isConsensusReportAlreadyProcessing {
		minSleep := time.Second * 12 * 32
		maxSleep := time.Second * 12 * 32 * 2
		return errs.NewSleepError("Consensus Report Already Processing (Main data already submitted)", timetool.RandomTime(minSleep, maxSleep))
	}

	if dataHash != memberInfo.CurrentFrameMemberReport {
		dataHashStr := hexutil.Encode(dataHash[:])

		err := v.submitReport(ctx, dataHash, refSlot, consensusVersion)
		if err != nil {
			return errors.Wrap(err, "")
		}
		reportJson, _ := json.Marshal(reportData)
		logger.Info("Send reportData's hash.", zap.String("hash", dataHashStr), zap.String("ReportData", string(reportJson)))
	} else {
		logger.Info("Provided hash already submitted.")
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
	logger.Info("submit consensus Report success.",
		zap.String("tx hash", tx.Hash().String()),
		zap.String("refSlot", refSlot.String()),
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

	return &MemberInfo{
		IsReportMember:              memberConsensusState.IsMember,
		IsFastLane:                  memberConsensusState.IsFastLane,
		CanReport:                   memberConsensusState.CanReport,
		LastReportRefSlot:           memberConsensusState.LastMemberReportRefSlot,
		FastLaneLengthSlot:          frameConfig.FastLaneLengthSlots,
		CurrentFrameRefSlot:         currentFrame.RefSlot,
		DeadlineSlot:                currentFrame.ReportProcessingDeadlineSlot,
		CurrentFrameMemberReport:    memberConsensusState.CurrentFrameMemberReport,
		CurrentFrameConsensusReport: memberConsensusState.CurrentFrameConsensusReport,
	}, nil
}

func (v *HashConsensusHelper) GetRefSlotAndIsReport(ctx context.Context) (*big.Int, bool, error) {
	err := v.ReportContract.CheckContractVersions(ctx)
	if err != nil {
		return nil, false, errors.Wrap(err, "")
	}

	headSlot, err := consensus.ConsensusClient.CustomizeBeaconService.HeadSlot(ctx)
	if err != nil {
		return nil, false, errors.Wrap(err, "")
	}

	memberInfo, err := v.GetMemberInfo(ctx)
	if err != nil {
		return nil, false, errors.Wrap(err, "")
	}

	if !memberInfo.CanReport {
		logger.Debug("Member's ConsensusState is not CanReport.")
		return memberInfo.CurrentFrameRefSlot, false, nil
	}

	// Slot < CurrentFrameRefSlot
	if headSlot.Cmp(memberInfo.CurrentFrameRefSlot) == -1 {
		logger.Info("Reference slot is not yet finalized.",
			zap.String("headSlot", headSlot.String()),
			zap.String("CurrentFrameRefSlot", memberInfo.CurrentFrameRefSlot.String()))
		return memberInfo.CurrentFrameRefSlot, false, nil
	}

	// Slot > DeadlineSlot
	if headSlot.Cmp(memberInfo.DeadlineSlot) == 1 {
		logger.Info("Deadline missed.",
			zap.String("headSlot", headSlot.String()),
			zap.String("DeadlineSlot", memberInfo.DeadlineSlot.String()))
		return memberInfo.CurrentFrameRefSlot, false, nil
	}

	return memberInfo.CurrentFrameRefSlot, true, nil
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
