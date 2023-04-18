// description:
// @author renshiwei
// Date: 2022/10/6 14:36

package cmd

import (
	"context"
	"fmt"
	"github.com/NodeDAO/oracle-go/app"
	"github.com/NodeDAO/oracle-go/cmd/oracle"
	"github.com/NodeDAO/oracle-go/cmd/version"
	"github.com/NodeDAO/oracle-go/config"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"os"
)

var (
	cfgFile string
	cliName = config.Config.Cli.Name
)

var RootCmd = &cobra.Command{
	Use:          cliName,
	Short:        cliName,
	SilenceUsage: true,
	Long:         cliName + `:https://github.com/NodeDAO/oracle-go`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			tip()
			return errors.New("requires at least one arg")
		}
		return nil
	},
	PersistentPreRun: func(*cobra.Command, []string) {
		ctx := context.Background()
		// init
		app.InitServer(ctx)
	},
	Run: func(cmd *cobra.Command, args []string) {
		tip()
	},
}

func tip() {
	usageStr := `Welcome to use ` + cliName + `:` + ` use ` + `-h` + ` see cli`
	usageStr1 := `You can also refer to the related content of https://github.com/NodeDAO/oracle-go 的相关内容`
	fmt.Printf("%s\n", usageStr)
	fmt.Printf("%s\n", usageStr1)
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file path")

	RootCmd.AddCommand(version.StartCmd)
	RootCmd.AddCommand(oracle.OracleCmd)
}

// Execute : apply commands
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}

func initConfig() {
	config.InitConfig(cfgFile)
}
