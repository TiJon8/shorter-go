package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type AppConfig struct {
	Env string `envconfig:"ENV" default:"prod"`
	Addr string `envconfig:"HTTP_ADDR" default:":8080"`
	StoragePath string `envconfig:"STORAGE_PATH" requered:"true"`
	LogLevel string `envconfig:"LOG_LEVEL" default:"info"`
}


func newAppConfig() (*AppConfig, error) {
	var s AppConfig
	if err := envconfig.Process("", &s); err != nil {
		return nil, fmt.Errorf("Failed to procces environment variables")
	}
	return &s, nil
}


func AppConfigMust() *AppConfig {
	config, err := newAppConfig()
	if err != nil {
		panic("error with processing env")
	}
	return config
}