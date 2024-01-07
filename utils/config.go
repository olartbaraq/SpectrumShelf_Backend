package utils

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DBdriver              string        `mapstructure:"DB_DRIVER"`
	DBsource              string        `mapstructure:"DB_SOURCE"`
	DBsourceLive          string        `mapstructure:"DB_SOURCE_LIVE"`
	SigningKey            string        `mapstructure:"SIGNING_KEY"`
	AccessTokenExpiresIn  time.Duration `mapstructure:"ACCESS_TOKEN_EXPIRED_IN"`
	RefreshTokenExpiresIn time.Duration `mapstructure:"REFRESH_TOKEN_EXPIRED_IN"`
	AccessTokenMaxAge     int           `mapstructure:"ACCESS_TOKEN_MAXAGE"`
	RefreshTokenMaxAge    int           `mapstructure:"REFRESH_TOKEN_MAXAGE"`
}

func LoadDBConfig(path string) (config *Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

// func LoadOtherConfig(path string) (config2 *Config, err error) {
// 	viper.AddConfigPath(path)
// 	viper.SetConfigName("app")
// 	viper.SetConfigType("env")

// 	viper.AutomaticEnv()

// 	err = viper.ReadInConfig()
// 	if err != nil {
// 		return nil, err
// 	}

// 	err = viper.Unmarshal(&config2)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return config2, nil
// }
