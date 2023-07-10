// description:
// @author renshiwei
// Date: 2023/7/4

package largestake

import (
	"context"
	"fmt"
	"github.com/NodeDAO/oracle-go/app"
	"github.com/NodeDAO/oracle-go/common/errs"
	"github.com/NodeDAO/oracle-go/common/logger"
	"github.com/NodeDAO/oracle-go/config"
	"github.com/NodeDAO/oracle-go/consensus"
	"github.com/NodeDAO/oracle-go/contracts"
	consensusApi "github.com/attestantio/go-eth2-client/api/v1"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"math/big"
	"testing"
)

func TestProcessLargeStakeOracleReport(t *testing.T) {
	var err error
	ctx := context.Background()
	config.InitConfig("../../conf/config-goerli.dev.yaml")
	logger.InitLog("debug", "console")
	app.InitServer(ctx)

	frame, err := contracts.HashConsensusContract.Contract.GetCurrentFrame(nil)
	require.NoError(t, err)

	// executionBlock
	executionBlock, err := consensus.ConsensusClient.CustomizeBeaconService.ExecutionBlock(ctx, frame.RefSlot.String())
	require.NoError(t, err)

	helper := NewLargeStakeHelper(frame.RefSlot, big.NewInt(2), executionBlock.BlockNumber)

	largeStakeReportRes, err := helper.ProcessReportData(ctx)

	if sleepErr, ok := errors.Cause(err).(*errs.SleepError); ok {
		logger.Debug("[LargeStakeOracle] sleep",
			zap.String("msg", sleepErr.Msg),
			zap.String("sleep time", sleepErr.Sleep.String()),
		)
		err = nil
	}
	require.NoError(t, err)

	fmt.Printf("largeStakeReportRes:%+v\n", largeStakeReportRes)
}

func TestFilterExitedSlashedValidator(t *testing.T) {
	var err error
	ctx := context.Background()
	config.InitConfig("../../conf/config-goerli.dev.yaml")
	logger.InitLog("debug", "console")
	app.InitServer(ctx)

	frame, err := contracts.HashConsensusContract.Contract.GetCurrentFrame(nil)
	require.NoError(t, err)

	// executionBlock
	executionBlock, err := consensus.ConsensusClient.CustomizeBeaconService.ExecutionBlock(ctx, frame.RefSlot.String())
	require.NoError(t, err)

	helper := NewLargeStakeHelper(frame.RefSlot, big.NewInt(2), executionBlock.BlockNumber)

	validatorExaMap := make(map[string]*LargeStakeValidator)

	pubkeys := []string{"0x81e9238fa1a672954d6c07d6efd7bf42c8b1827347da2bb73b5921065f8857fa1867b2a3ce4677c97ecbbad0b1d5299f"}

	validators, err := consensus.ConsensusClient.CustomizeBeaconService.ValidatorsByPubKey(ctx, frame.RefSlot.String(), pubkeys)

	for _, pubkey := range pubkeys {
		validator, ok := validators[pubkey]
		// Handle pubkeys that are not query by Beacon
		// set balance = 32 GWEI
		if !ok {
			validator = &consensusApi.Validator{
				Balance: 32e9,
			}
		}
		validatorExaMap[pubkey] = &LargeStakeValidator{
			StakingId: big.NewInt(1),
		}
		validatorExaMap[pubkey].Validator = validator
	}

	validatorInfos, slashInfos, err := helper.filterExitedSlashedValidator(ctx, validatorExaMap)
	fmt.Printf("Validator:%+v\n", validatorInfos)
	fmt.Printf("slashInfos:%+v\n", slashInfos)
}
