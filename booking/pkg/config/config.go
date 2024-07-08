package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Environment      string `mapstructure:"ENV"`
	DatabaseHost     string `mapstructure:"DATABASE_HOST"`
	DatabasePassword string `mapstructure:"DATABASE_PASSWORD"`
	DatabaseUsername string `mapstructure:"DATABASE_USERNAME"`
	DatabasePort     int    `mapstructure:"DATABASE_PORT"`
	DatabaseName     string `mapstructure:"DATABASE_NAME"`
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) Load() error {
	if viper.Get("ENV") == "production" {
		return nil
	}

	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	err = viper.Unmarshal(c)

	if err != nil {
		return err
	}

	return nil
}
