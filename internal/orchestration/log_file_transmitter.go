package orchestration

import (
	"context"
	"errors"
	"github.com/jrolstad/log-analytics-platform/internal/clients"
	"github.com/jrolstad/log-analytics-platform/internal/config"
	"github.com/jrolstad/log-analytics-platform/internal/logging"
	"github.com/oracle/oci-go-sdk/v49/common"
	"github.com/oracle/oci-go-sdk/v49/objectstorage"
)

func PublishFilesInBuckets(appConfig *config.AppConfig) error {
	client, err := clients.GetObjectStorageClient(appConfig)
	if err != nil {
		return err
	}

	logging.LogEvent("Publishing bucket files",
		"region", appConfig.Region, "buckets", appConfig.Buckets, "directories", appConfig.Directories)

	processErrors := make([]error, 0)
	successfulCount := 0
	for _, bucketName := range appConfig.Buckets {
		err := processBucket(bucketName, appConfig.Directories, appConfig.Namespace, client)
		if err != nil {
			processErrors = append(processErrors, err)
		} else {
			successfulCount++
		}
	}

	logging.LogEvent("Published bucket files",
		"region", appConfig.Region, "errors", len(processErrors), "success", successfulCount)

	return errors.Join(processErrors...)
}

func processBucket(bucketName string,
	directories []string,
	namespace string,
	client objectstorage.ObjectStorageClient) error {
	bucket, err := getBucket(bucketName, namespace, client)
	if err != nil {
		return err
	}

	logging.LogEvent("Processing bucket",
		"name", bucket.Name, "id", bucket.Id, "namespace", bucket.Namespace)

	processErrors := make([]error, 0)
	for _, directory := range directories {
		err := processDirectory(&bucket, directory, client)
		if err != nil {
			processErrors = append(processErrors, err)
		}
	}

	return errors.Join(processErrors...)
}

func processDirectory(bucket *objectstorage.Bucket,
	directory string,
	client objectstorage.ObjectStorageClient) error {
	bucketObjects, err := listObjectsInBucket(bucket, directory, client)
	if err != nil {
		return err
	}

	logging.LogEvent("Processing bucket directory",
		"bucket", bucket.Name,
		"directory", directory,
		"fileCount", len(bucketObjects))

	err = publishBucketObjects(bucket, bucketObjects)
	if err != nil {
		return err
	}

	return nil
}

func listObjectsInBucket(bucket *objectstorage.Bucket,
	directory string,
	client objectstorage.ObjectStorageClient) ([]objectstorage.ObjectSummary, error) {
	result := make([]objectstorage.ObjectSummary, 0)
	listErrors := make([]error, 0)

	hasNext := true
	var nextPage *string

	for hasNext == true {
		request := objectstorage.ListObjectsRequest{
			NamespaceName: bucket.Namespace,
			BucketName:    bucket.Name,
			Prefix:        common.String(directory),
		}
		if nextPage != nil {
			request.StartAfter = nextPage
		}

		response, err := client.ListObjects(context.Background(), request)
		if err != nil {
			listErrors = append(listErrors, err)
		}

		result = append(result, response.Objects...)

		nextPage = response.NextStartWith
		hasNext = nextPage != nil
	}

	return result, errors.Join(listErrors...)
}

func getBucket(bucketName string,
	namespace string,
	client objectstorage.ObjectStorageClient) (objectstorage.Bucket, error) {
	request := objectstorage.GetBucketRequest{
		NamespaceName: common.String(namespace),
		BucketName:    common.String(bucketName),
	}
	bucket, err := client.GetBucket(context.Background(), request)
	if err != nil {
		return objectstorage.Bucket{}, err
	}
	return bucket.Bucket, nil
}

func publishBucketObjects(bucket *objectstorage.Bucket, toPublish []objectstorage.ObjectSummary) error {
	for _, item := range toPublish {
		logging.LogEvent("Publishing bucket object",
			"name", item.Name,
			"bucket", bucket.Name,
			"namespace", bucket.Namespace)
	}
	return nil
}
