// desc:
// @author renshiwei
// Date: 2023/4/12 11:09

package withdraw

import (
	"context"
	"github.com/NodeDAO/oracle-go/app"
	"github.com/NodeDAO/oracle-go/app/withdraw"
	"github.com/NodeDAO/oracle-go/common/errs"
	"github.com/NodeDAO/oracle-go/common/logger"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"runtime"
	"time"
)

var (
	WithdrawOracleCmd = &cobra.Command{
		Use:   "withdraw",
		Short: "Run Withdraw Oracle Server",
		Long:  `Run Withdraw Oracle Server`,
		PreRun: func(*cobra.Command, []string) {
			ctx := context.Background()
			app.InitServer(ctx)
		},
		Run: func(cmd *cobra.Command, args []string) {
			for {
				run()
			}
		},
	}
)

func run() {
	defer func() {
		if err := recover(); err != nil {
			switch err.(type) {
			case runtime.Error:
				logger.Errorf("runtime error:%+v", err)
				withdraw.DefaultRandomSleep()
			default:
				logger.Errorf("error::%+v", err)
			}
		}
	}()

	ctx := context.Background()
	w := new(withdraw.WithdrawHelper)

	err := w.ProcessReport(ctx)
	if err != nil {
		switch errors.Cause(err).(type) {
		case *errs.SleepError:
			// If err is of type Sleep Error, the oracle program has reached the condition that needs to be slept and re-executed
			if sleepErr, ok := errors.Cause(err).(*errs.SleepError); ok {
				logger.Debug("withdraw oracle sleep",
					zap.String("msg", sleepErr.Msg),
					zap.String("sleep time", sleepErr.Sleep.String()),
				)
				time.Sleep(sleepErr.Sleep)
			}
		default:
			logger.Errorf("err:%+v", err)
			withdraw.DefaultRandomSleep()
		}
		return
	} else {
		sleepTime := withdraw.RandomSleepTime()

		logger.Debug("withdraw oracle sleep success run once for sleep.",
			zap.String("sleep time", sleepTime.String()),
		)
		time.Sleep(sleepTime)
	}
}
