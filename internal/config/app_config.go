package config

import (
	"os"
	"strings"
)

func GetAppConfig() *AppConfig {
	return &AppConfig{
		Region:      os.Getenv("log_analyzer_region"),
		Buckets:     strings.Split(os.Getenv("log_analzyer_buckets"), ","),
		Directories: strings.Split(os.Getenv("log_analzyer_directories"), ","),
	}
}

type AppConfig struct {
	Region      string
	Buckets     []string
	Directories []string
}
