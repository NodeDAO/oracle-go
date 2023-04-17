// description:
// @author renshiwei
// Date: 2022/10/5 17:33

package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"strings"
)

func InitConfig(configDir string) {
	viper.SetConfigFile(configDir)
	viper.SetConfigType("yaml")
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.AutomaticEnv() // read in environment variables that match

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	//viper.MergeInConfig()

	viper.WatchConfig()

	if err := viper.Unmarshal(&Config); err != nil {
		panic(fmt.Errorf("read config file err. file: %s", configDir))
	}

	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := viper.Unmarshal(&Config); err != nil {
			fmt.Println(err)
		}
	})

}
