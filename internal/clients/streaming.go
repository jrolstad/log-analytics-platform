package clients

import (
	"github.com/jrolstad/log-analytics-platform/internal/config"
	"github.com/oracle/oci-go-sdk/v49/common"
	"github.com/oracle/oci-go-sdk/v49/streaming"
)

func GetStreamingClient(appConfig *config.AppConfig) (streaming.StreamClient, error) {
	configProvider := common.NewRawConfigurationProvider(appConfig.TenancyOcid,
		appConfig.UserOcid,
		appConfig.Region,
		appConfig.Fingerprint,
		appConfig.PrivateKey,
		nil)

	client, err := streaming.NewStreamClientWithConfigurationProvider(configProvider, appConfig.FileStreamEndpoint)
	if err != nil {
		return streaming.StreamClient{}, err
	}

	return client, err
}
