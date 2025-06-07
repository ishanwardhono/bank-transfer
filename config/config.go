package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	ServerHost string
	ServerPort string
	LogLevel   string
}

func LoadConfig() (*Config, error) {
	v := viper.New()
	v.SetConfigFile(".env")
	v.SetConfigType("env")

	// Read the config file
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("failed to read config file: %w", err)
		}
	}

	return &Config{
		ServerHost: v.GetString("SERVER_HOST"),
		ServerPort: v.GetString("SERVER_PORT"),
		LogLevel:   v.GetString("LOG_LEVEL"),
	}, nil
}

func (c *Config) GetServerAddress() string {
	return fmt.Sprintf("%s:%s", c.ServerHost, c.ServerPort)
}
