package config

import "github.com/kelseyhightower/envconfig"

const envPrefix = "QA"

type Config struct {
	Host   string `split_words:"true" default:"localhost:8001"`
	DbHost string `split_words:"true" default:"localhost"`
	DbPort int    `split_words:"true" default:"81"`
}

func FromEnv() (*Config, error) {
	cfg := &Config{}
	err := envconfig.Process(envPrefix, cfg)
	return cfg, err
}
