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
	"github.com/NodeDAO/oracle-go/contracts"
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

	helper := NewLargeStakeHelper(frame.RefSlot, big.NewInt(2))

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
