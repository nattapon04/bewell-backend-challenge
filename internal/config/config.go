package config

import (
	"bewell-backend-challenge/util/helpers/logger"

	"github.com/spf13/viper"
)

type Config struct {
	AppName string `mapstructure:"APPLICATION_NAME" default:"redemtion-api"`
	AppEnv  string `mapstructure:"APP_ENV"`
	BaseURL string `mapstructure:"BASE_URL" default:"http://localhost"`
	Version string `mapstructure:"VERSION" default:"1.0.0"`
	AppPort string `mapstructure:"APP_PORT" default:"8080"`
}

func Read() *Config {
	conf := Config{}

	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log(conf.AppName, err)
	}

	if err := viper.Unmarshal(&conf); err != nil {
		log(conf.AppName, err)
	}

	return &conf
}

func log(appName string, err error) {
	logData := logger.SetFormatter(appName, 0, `.env`, nil)
	logger.New().WithFields(logData).Error(err)
}
