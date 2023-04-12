// desc:
// @author renshiwei
// Date: 2023/4/12 10:41

package oracle

import (
	"github.com/NodeDAO/oracle-go/cmd/oracle/withdraw"
	"github.com/spf13/cobra"
)

var (
	OracleCmd = &cobra.Command{
		Use:   "oracle",
		Short: "Run Oracle Server",
		Long:  `Run Oracle Server.`,
	}
)

func init() {
	OracleCmd.AddCommand(withdraw.WithdrawOracleCmd)
}
