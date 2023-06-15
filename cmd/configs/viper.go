package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Port      int    `mapstructure:"PORT"`
	SecretKey string `mapstructure:"SECRET_KEY"`
}

func ReadConfig() (Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs") // Set the directory path where the config file is located

	// Read the configuration file
	err := viper.ReadInConfig() // Read the config file and handle any errors
	if err != nil {
		return Config{}, fmt.Errorf("failed to read config file: %w", err)
	}

	var config Config
	err = viper.Unmarshal(&config) // Unmarshal the config into the struct
	if err != nil {
		return Config{}, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return config, nil
}