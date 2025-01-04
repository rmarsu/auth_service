package config

import (
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Enviroment string     `yaml:"env" env-default:"local"`
	Database   Database   `yaml:"database"`
	GRPC       GRPCConfig `yaml:"grpc"`
}

type Database struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type GRPCConfig struct {
	Port              int `yaml:"port"`
	MaxConnectionAge  time.Duration `yaml:"max_connection_age"`
	MaxConnectionIdle time.Duration `yaml:"max_connection_idle"`
	Timeout           time.Duration `yaml:"timeout"`
}

func MustLoad(cfgPath string) *Config {
	data, err := os.ReadFile(cfgPath)
	if err != nil {
		panic(err)
	}
	var config *Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		panic(err)
	}
	return config

}
