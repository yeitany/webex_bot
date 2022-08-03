package utils

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	RoomID  string `mapstructure:"ROOM_ID"`
	Token   string `mapstructure:"WEBEX_TOKEN"`
	BaseUrl string `mapstructure:"WEBEX_BASE_URL"`
}

func NewConfig() Config {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatal(err)
	}

	return config
}
