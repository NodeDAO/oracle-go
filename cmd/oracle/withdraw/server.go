// desc:
// @author renshiwei
// Date: 2023/4/12 11:09

package withdraw

import (
	"context"
	"github.com/NodeDAO/oracle-go/app/withdraw"
	"github.com/NodeDAO/oracle-go/common/logger"
	"github.com/spf13/cobra"
)

var (
	WithdrawOracleCmd = &cobra.Command{
		Use:   "withdraw",
		Short: "Run Withdraw Oracle Server",
		Long:  `Run Withdraw Oracle Server`,
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
)

func run() {
	ctx := context.Background()
	w := new(withdraw.WithdrawHelper)

	//todo for
	//for {
	err := w.ProcessReport(ctx)
	if err != nil {
		logger.Errorf("err:%+v", err)
	}
	//}
}
