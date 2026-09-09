package config

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
)


type ServerConfig struct {
	Addr string `envconfig:"ADDR" required:"true"`
	GracefullShutdownDuration time.Duration `envconfig:"GRACEFULL_SHUTDOWNN_DURATION"`
}


func processServerConfig() (*ServerConfig, error) {
	var config ServerConfig
	if err := envconfig.Process("http", &config); err != nil {
		return nil, fmt.Errorf("Error with processing server config: %w", err)
	}
	return &config, nil
}

func ServerConfigMust() *ServerConfig {
	cfg, err := processServerConfig()
	if err != nil {
		panic(err)
	}
	return cfg
}