// description:
// @author renshiwei
// Date: 2022/10/5 17:06

package config

//go:generate go-bindata -pkg=config -nocompress -o=default_conf.go ../conf/config-default.yaml

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
		Network    string
		ElAddr     string
		ClAddr     string
		PrivateKey string
	}

	Oracle struct {
		IsReportData          bool
		IsSimulatedReportData bool
	}
}
