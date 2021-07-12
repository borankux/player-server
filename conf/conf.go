package conf

import (
	"github.com/spf13/viper"
	"log"
)

var conf *viper.Viper

func Init(env string) {
	conf = viper.GetViper()
	conf.AddConfigPath("conf")
	conf.SetConfigName(env)
	err:= conf.ReadInConfig()
	if err != nil {
		log.Fatalf("Failed to load the config file %v", err)
	}
}

func GetConf() * viper.Viper{
	return conf
}
