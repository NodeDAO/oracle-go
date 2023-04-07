package version

import (
	"fmt"
	"github.com/NodeDAO/oracle-go/config/global"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	StartCmd = &cobra.Command{
		Use:     "version",
		Short:   "Get version info",
		Example: global.Config.Cli.Name + " version",
		PreRun: func(cmd *cobra.Command, args []string) {

		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func run() error {
	fmt.Printf(global.Config.Cli.Name+" version: %s\n", color.GreenString(global.Config.Server.Version))
	return nil
}
