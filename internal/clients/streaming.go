package clients

import (
	"github.com/jrolstad/log-analytics-platform/internal/config"
	"github.com/oracle/oci-go-sdk/v49/streaming"
)

func GetStreamingClient(appConfig *config.AppConfig) (streaming.StreamClient, error) {
	configProvider := getConfigurationProvider(appConfig)

	client, err := streaming.NewStreamClientWithConfigurationProvider(configProvider, appConfig.FileStreamEndpoint)
	if err != nil {
		return streaming.StreamClient{}, err
	}

	return client, err
}
