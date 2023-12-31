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
		Namespace:   os.Getenv("log_analyzer_namespace"),

		TenancyOcid: os.Getenv("log_analyzer_tenancyocid"),
		UserOcid:    os.Getenv("log_analyzer_userocid"),
		Fingerprint: os.Getenv("log_analyzer_fingerprint"),
		PrivateKey:  os.Getenv("log_analyzer_privatekey"),

		FileStreamEndpoint: os.Getenv("log_analyzer_filestream_endpoint"),
		FileStreamId:       os.Getenv("log_analyzer_filestream_id"),
	}
}

type AppConfig struct {
	Region      string
	Namespace   string
	Buckets     []string
	Directories []string

	TenancyOcid string
	UserOcid    string
	Fingerprint string
	PrivateKey  string

	FileStreamEndpoint string
	FileStreamId       string
}
