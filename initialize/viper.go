package initialize

import (
	"example.com/m/v2/config"
	"example.com/m/v2/global"
	"fmt"
	"github.com/spf13/viper"
)

func InitConfig() {
	v := viper.New()
	// 设置配置文件
	v.SetConfigFile("./config.yaml")
	v.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	serveConfig := config.Serve{}
	err = v.Unmarshal(serveConfig)
	if err != nil {
		fmt.Println(err)
	}

	global.CONFIG = serveConfig
}
