// description:
// @author renshiwei
// Date: 2022/10/5 17:57

package app

import (
	"context"
	"github.com/NodeDAO/oracle-go/common/logger"
	"github.com/NodeDAO/oracle-go/config"
	"github.com/NodeDAO/oracle-go/consensus"
	"github.com/NodeDAO/oracle-go/eth1"
	"sync"
	"time"
)

var once sync.Once

func InitServer(ctx context.Context, configDir string) {
	timeout := 1 * time.Minute
	var err error

	once.Do(func() {
		config.InitConfig(configDir)
		logger.InitLog()
		consensus.ConsensusClient, err = consensus.New(ctx, config.Config.Eth.ClAddr, timeout)
		if err != nil {
			logger.Errorf("err:%+v", err)
			panic(err)
		}

		eth1.ElClient, err = eth1.NewEthClient(ctx, config.Config.Eth.ElAddr)
	})
}
