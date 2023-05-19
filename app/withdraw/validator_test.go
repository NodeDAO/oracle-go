package withdraw

import (
	"context"
	"fmt"
	"github.com/NodeDAO/oracle-go/app"
	"github.com/NodeDAO/oracle-go/common/logger"
	"github.com/NodeDAO/oracle-go/config"
	"github.com/NodeDAO/oracle-go/consensus"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"math/big"
	"testing"
)

func TestValidatorBalance(t *testing.T) {
	var err error
	ctx := context.Background()
	config.InitConfig("../../conf/config-mainnet.dev.yaml")
	logger.InitLog("debug", "console")
	app.InitServer(ctx)

	w := new(WithdrawHelper)
	w.clBalance = big.NewInt(0)

	consensusContract, err := w.oracle.GetConsensusContract(ctx)
	require.NoError(t, err)

	currentFrame, err := consensusContract.GetCurrentFrame(nil)
	require.NoError(t, err)
	logger.Debug("current frame",
		zap.String("RefSlot", currentFrame.RefSlot.String()),
		zap.String("DeadlineSlot", currentFrame.ReportProcessingDeadlineSlot.String()),
	)

	headSlot, err := consensus.ConsensusClient.CustomizeBeaconService.HeadSlot(ctx)
	if headSlot.Cmp(currentFrame.RefSlot) == -1 || headSlot.Cmp(currentFrame.ReportProcessingDeadlineSlot) == 1 {
		panic(fmt.Sprintf("slot is not correct. headslot:%s", headSlot.String()))
	}

	w.refSlot = currentFrame.RefSlot

	err = w.obtainValidatorConsensusInfo(ctx)
	require.NoError(t, err)

	logger.Debug("-------start calculate validator balance-----", zap.Int("count", len(w.validatorExaMap)))
	for pubkey, exa := range w.validatorExaMap {
		// sum cl balance
		w.clBalance = new(big.Int).Add(w.clBalance, big.NewInt(int64(exa.Validator.Balance)))

		fmt.Printf("pubkey:%s balance:%v status:%s \n", pubkey, exa.Validator.Balance, exa.Validator.Status)
	}
	logger.Debug("-------end calculate validator balance-----")

	logger.Debugf("clBalance:%s", w.clBalance.String())
}
