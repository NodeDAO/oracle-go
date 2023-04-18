// desc:
// @author renshiwei
// Date: 2023/4/12 11:09

package withdraw

import (
	"context"
	"github.com/NodeDAO/oracle-go/app/withdraw"
	"github.com/NodeDAO/oracle-go/common/logger"
	"github.com/spf13/cobra"
	"runtime"
)

var (
	WithdrawOracleCmd = &cobra.Command{
		Use:   "withdraw",
		Short: "Run Withdraw Oracle Server",
		Long:  `Run Withdraw Oracle Server`,
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
		logger.Errorf("err:%+v", err)
	}

}
