package version

import (
	"fmt"
	"github.com/NodeDAO/oracle-go/config"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	StartCmd = &cobra.Command{
		Use:     "version",
		Short:   "Get version info",
		Example: config.Config.Cli.Name + " version",
		PreRun: func(cmd *cobra.Command, args []string) {

		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func run() error {
	fmt.Printf(config.Config.Cli.Name+" version: %s\n", color.GreenString(config.Config.Server.Version))
	return nil
}
