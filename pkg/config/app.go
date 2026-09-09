package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type AppConfig struct {
	Env string `envconfig:"ENV" default:"prod"`
	StoragePath string `envconfig:"STORAGE_PATH" required:"true"`
	LogLevel string `envconfig:"LOG_LEVEL" default:"info"`
}


func newAppConfig() (*AppConfig, error) {
	var s AppConfig
	if err := envconfig.Process("", &s); err != nil {
		return nil, fmt.Errorf("Failed to process environment variables: %v", err)
	}
	return &s, nil
}


func AppConfigMust() *AppConfig {
	config, err := newAppConfig()
	if err != nil {
		panic(err)
	}
	return config
}