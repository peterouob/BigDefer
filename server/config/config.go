package config

import (
	"os"

	"github.com/spf13/viper"
)

var Config *viper.Viper

func init() {
	Config = viper.New()
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	Config.AddConfigPath(wd)
	Config.SetConfigName("config")
	Config.SetConfigType("yaml")

	if err := Config.ReadInConfig(); err != nil {
		panic(err)
	}
}
