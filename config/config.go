package config

import (
	"log"

	"github.com/spf13/viper"
)

var configName = "global"

func init() {
	viper.SetConfigName(configName) // name of config file (without extension)
	viper.SetConfigType("yaml")     //REQUIRED if the config file does not have the extension in the name

	viper.AddConfigPath(".") // 设置配置文件和可执行二进制文件在用一个目录
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			log.Fatal("no such config file")
		} else {
			// Config file was found but another error was produced
			log.Fatal("read config error")
		}
		log.Fatal(err) // 读取配置文件失败致命错误
	}

}
