package clients

import (
	"github.com/jrolstad/log-analytics-platform/internal/config"
	"github.com/oracle/oci-go-sdk/v49/common"
	"github.com/oracle/oci-go-sdk/v49/objectstorage"
)

func GetObjectStorageClient(appConfig *config.AppConfig) (objectstorage.ObjectStorageClient, error) {
	configProvider := common.NewRawConfigurationProvider(appConfig.TenancyOcid,
		appConfig.UserOcid,
		appConfig.Region,
		appConfig.Fingerprint,
		appConfig.PrivateKey,
		nil)

	client, err := objectstorage.NewObjectStorageClientWithConfigurationProvider(configProvider)
	if err != nil {
		return objectstorage.ObjectStorageClient{}, err
	}
	return client, nil
}
