// description:
// @author renshiwei
// Date: 2022/10/5 17:57

package config

import (
	"github.com/NodeDAO/oracle-go/common/logger"
	"sync"
)

var once sync.Once

func InitServer(configDir string) {
	once.Do(func() {
		InitConfig(configDir)
		logger.InitLog()
	})
}
