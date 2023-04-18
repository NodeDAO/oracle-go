// description:
// @author renshiwei
// Date: 2022/10/5 17:33

package config

import (
	"bytes"
	"fmt"
	"github.com/NodeDAO/oracle-go/conf"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"strings"
)

func InitConfig(configFile ...string) {
	//viper.SetConfigFile(defaultConfig)
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.AutomaticEnv() // read in environment variables that match

	defaultConfig, err := conf.Asset("conf/config-default.yaml")
	if err != nil {
		panic(fmt.Errorf("conf.Asset default config file: %s err:%+v\n", configFile, err))
	}

	if err := viper.ReadConfig(bytes.NewBuffer(defaultConfig)); err != nil {
		panic(fmt.Errorf("Fatal error default config file: %s err:%+v\n", defaultConfig, err))
	}

	if len(configFile) > 0 {
		for _, s := range configFile {
			if s != "" {
				viper.SetConfigFile(s)
			}
		}
	}

	err = viper.MergeInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s err:%+v\n", configFile, err))
	}

	viper.WatchConfig()

	if err := viper.Unmarshal(&Config); err != nil {
		panic(fmt.Errorf("read config file err. file: %s %s", defaultConfig, configFile))
	}

	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := viper.Unmarshal(&Config); err != nil {
			fmt.Println(err)
		}
	})

}
