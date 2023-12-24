package clients

import (
	"github.com/jrolstad/log-analytics-platform/internal/config"
	"github.com/oracle/oci-go-sdk/v49/objectstorage"
)

func GetObjectStorageClient(appConfig *config.AppConfig) (objectstorage.ObjectStorageClient, error) {
	configProvider := getConfigurationProvider(appConfig)

	client, err := objectstorage.NewObjectStorageClientWithConfigurationProvider(configProvider)
	if err != nil {
		return objectstorage.ObjectStorageClient{}, err
	}
	return client, nil
}
