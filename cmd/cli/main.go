package main

import (
	"github.com/jrolstad/log-analytics-platform/internal/config"
	"github.com/jrolstad/log-analytics-platform/internal/orchestration"
)

func main() {
	appConfig := config.GetAppConfig()

	err := orchestration.PublishFilesInBuckets(appConfig)
	if err != nil {
		panic(err)
	}
}
