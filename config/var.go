// description:
// @author renshiwei
// Date: 2022/10/5 17:06

//go:generate go-bindata -pkg=config -nocompress -o=default_conf.go ../conf/config-default.yaml

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
		Format string
	}

	Eth struct {
		Network    string
		ElAddr     string
		ClAddr     string
		PrivateKey string
	}

	Oracle struct {
		IsReportData                   bool
		IsSimulatedReportData          bool
		IsDifferentConsensusHashReport bool
	}
}
