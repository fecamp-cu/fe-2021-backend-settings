package configs

import (
	"github.com/spf13/viper"
	"log"
)

type Configuration struct {
	App      AppConfigs
	Postgres PostgresConfigs
	Redis    RedisConfigs
}

func LoadConfigs() *Configuration {
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
	return &configuration
}
