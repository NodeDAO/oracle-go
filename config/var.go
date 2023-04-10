// description:
// @author renshiwei
// Date: 2022/10/5 17:06

package config

var (
	Config ConfigYaml
)

type ConfigYaml struct {
	Server struct {
		Name    string
		Version string
		Http    struct {
			Host string
			Port int
		}
	}

	Cli struct {
		Name string
	}

	Log struct {
		Level struct {
			Server string
		}
	}

	Eth struct {
		ElAddr string
		ClAddr string
	}

	Contracts struct {
		Vnft           string
		Liq            string
		WithdrawOracle string
		HashConsensus  string
		ClVault        string
	}
}
