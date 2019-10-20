package cfg

import (
	"github.com/spf13/viper"
	"log"
)

func Init(pathToConfig string) {
	viper.SetConfigFile(pathToConfig)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Unable to read config: ", err)
	}
}

var (
	GetInt = viper.GetInt
)
