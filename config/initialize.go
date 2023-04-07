// description:
// @author renshiwei
// Date: 2022/10/5 17:57

package config

import (
	"context"
	"github.com/NodeDAO/oracle-go/common/logger"
	"github.com/NodeDAO/oracle-go/config/global"
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
		InitConfig(configDir)
		logger.InitLog()
		global.ConsensusClient, err = consensus.New(ctx, global.Config.Eth.ClAddr, timeout)
		if err != nil {
			logger.Errorf("err:%+v", err)
			panic(err)
		}

		global.ElClient, err = eth1.NewEthClient(ctx, global.Config.Eth.ElAddr)
	})
}
