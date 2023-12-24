package clients

import (
	"github.com/jrolstad/log-analytics-platform/internal/config"
	"github.com/oracle/oci-go-sdk/v49/common"
)

func getConfigurationProvider(appConfig *config.AppConfig) common.ConfigurationProvider {
	configProvider := common.NewRawConfigurationProvider(appConfig.TenancyOcid,
		appConfig.UserOcid,
		appConfig.Region,
		appConfig.Fingerprint,
		appConfig.PrivateKey,
		nil)

	return configProvider
}
