package config

import (
	"github.com/spf13/viper"
)

type Configuration struct {
	Port           string `mapstructure:"PORT"`
	DBHost         string `mapstructure:"POSTGRES_HOST"`
	DBUserName     string `mapstructure:"POSTGRES_USER"`
	DBUserPassword string `mapstructure:"POSTGRES_PASSWORD"`
	DBName         string `mapstructure:"POSTGRES_DB"`
	DBPort         string `mapstructure:"POSTGRES_PORT"`
	JWTSecretKey   string `mapstructure:"JWT_SECRET_KEY"`
}

var Config Configuration

func (config *Configuration) LoadConfig(path string) (err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(config)
	return
}
