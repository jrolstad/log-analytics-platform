package main

import (
	"github.com/jrolstad/log-analytics-platform/internal/clients"
	"github.com/jrolstad/log-analytics-platform/internal/config"
	"github.com/jrolstad/log-analytics-platform/internal/orchestration"
)

func main() {
	appConfig := config.GetAppConfig()

	objectStorageClient, err := clients.GetObjectStorageClient(appConfig)
	if err != nil {
		panic(err)
	}

	streamClient, err := clients.GetStreamingClient(appConfig)
	if err != nil {
		panic(err)
	}

	err = orchestration.PublishFilesInBuckets(appConfig, objectStorageClient, streamClient)
	if err != nil {
		panic(err)
	}
}
