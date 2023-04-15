// desc:
// @author renshiwei
// Date: 2023/4/11 15:28

package consensusModule

import (
	"context"
	"github.com/NodeDAO/oracle-go/common/logger"
	"github.com/NodeDAO/oracle-go/consensus"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"math/big"
)

func (v *HashConsensusHelper) ProcessReportHash(ctx context.Context, dataHash [32]byte, refSlot, consensusVersion *big.Int) error {
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

	if dataHash != memberInfo.CurrentFrameMemberReport {
		logger.Infof("Send report hash. hash:%s", dataHash)
		err := v.submitReport(ctx, dataHash, refSlot, consensusVersion)
		if err != nil {
			return errors.Wrap(err, "")
		}
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

	tx, err := consensusContract.SubmitReport(v.KeyTransactOpts, refSlot, dataHash, consensusVersion)
	if err != nil {
		return errors.Wrapf(err, "Failed to submit consensus report. refslot:%s, consensusVersion:%s", refSlot.String(), consensusVersion.String())
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

	memberConsensusState, err := consensusContract.GetConsensusStateForMember(nil, v.KeyTransactOpts.From)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to get HashConsensus's GetConsensusStateForMember.")
	}

	return &MemberInfo{
		IsReportMember:              memberConsensusState.IsMember,
		IsFastLane:                  memberConsensusState.IsFastLane,
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
