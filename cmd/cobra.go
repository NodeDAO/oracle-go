// description:
// @author renshiwei
// Date: 2022/10/6 14:36

package cmd

import (
	"fmt"
	"github.com/NodeDAO/oracle-go/cmd/version"
	"github.com/NodeDAO/oracle-go/common/global"
	"github.com/NodeDAO/oracle-go/config"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"os"
)

var (
	cfgFile string
	cliName = global.Config.Cli.Name
)

var rootCmd = &cobra.Command{
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
	PersistentPreRunE: func(*cobra.Command, []string) error { return nil },
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
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "conf/config-dev.yaml", "config file (default is config/config.yaml)")
	// init config file ...
	config.InitServer(cfgFile)

	rootCmd.AddCommand(version.StartCmd)
}

// Execute : apply commands
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
