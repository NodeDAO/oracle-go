package config

import (
	"fmt"
	"testing"
)

func TestConfigMerge(t *testing.T) {
	InitConfig("../conf/config-dev.yaml")
	fmt.Println("server.name", Config.Server.Name)
	fmt.Println("eth.network", Config.Eth.Network)
	fmt.Println("cli.name", Config.Cli.Name)
}
