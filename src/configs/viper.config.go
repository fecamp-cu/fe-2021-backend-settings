package configs

import (
	"log"
	"sync"

	"github.com/spf13/viper"
)

type Configuration struct {
	App      AppConfigs
	Postgres PostgresConfigs
	Redis    RedisConfigs
}

var configs Configuration
var lock sync.Once

func initViper() {
	viper.SetConfigName("config")
	viper.AddConfigPath("./configs")
	viper.AddConfigPath("../configs")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("fatal error config file: %s \n", err)
	}

	var configuration Configuration
	if err := viper.Unmarshal(&configuration); err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	configs = configuration
}

func GetConfigs() Configuration {
	lock.Do(initViper)
	return configs
}
