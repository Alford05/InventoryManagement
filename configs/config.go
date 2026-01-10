package configs

import (
	"os"

	"gopkg.in/yaml.v3"
)

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

type Config struct {
	Server struct {
		Port int `yaml:"port"`
	}
	Database DatabaseConfig
}

func LoadConfig() *Config {
	cfg := &Config{}

	file, err := os.ReadFile("configs/config.yaml")
	if err == nil {
		_ = yaml.Unmarshal(file, cfg)
	}
	if v := os.Getenv("DB_HOST"); v != "" {
		cfg.Database.Host = v
	}
	if v := os.Getenv("DB_USER"); v != "" {
		cfg.Database.User = v
	}
	if v := os.Getenv("DB_PASSWORD"); v != "" {
		cfg.Database.Password = v
	}
	if v := os.Getenv("DB_NAME"); v != "" {
		cfg.Database.Name = v
	}

	return cfg
}
