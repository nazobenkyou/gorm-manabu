package config

import "github.com/kelseyhightower/envconfig"

func LoadConfig() {
	_ = envconfig.Process("", &AppConfig)
	_ = envconfig.Process("", &Database)
}
