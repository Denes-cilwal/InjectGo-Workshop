package lib

import (
	"github.com/spf13/viper"
	"log"
)

type Env struct {
	ServerPort string `mapstructure:"SERVER_PORT"`
	DBUsername string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASS"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBName     string `mapstructure:"DB_NAME"`
	DBType     string `mapstructure:"DB_TYPE"`
}

var globalEnv = Env{}

func NewEnv() *Env {
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("cannot read configuration", err)
	}

	err = viper.Unmarshal(&globalEnv)
	if err != nil {
		log.Fatal("environment cant be loaded: ", err)
	}

	return &globalEnv
}
