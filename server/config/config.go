package config

import (
	"github.com/spf13/viper"
)

var Conf *viper.Viper

func init() {
	Conf = viper.New()
	Conf.AddConfigPath(".")
	Conf.SetConfigFile("config.yaml")

	if err := Conf.ReadInConfig(); err != nil {
		panic(err)
	}
}
