// description:
// @author renshiwei
// Date: 2022/10/5 17:06

package global

import (
	"github.com/NodeDAO/oracle-go/consensus"
	"github.com/NodeDAO/oracle-go/eth1"
)

var (
	Config          ConfigYaml
	ConsensusClient *consensus.Consensus
	ElClient        *eth1.EthClient
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
	}
}
