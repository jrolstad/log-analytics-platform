package mapping

import (
	"github.com/jrolstad/log-analytics-platform/internal/models"
	"github.com/oracle/oci-go-sdk/v49/objectstorage"
)

func MapToFileMessage(bucket *objectstorage.Bucket, summary objectstorage.ObjectSummary) models.FilePublished {
	return models.FilePublished{
		BucketName:      *bucket.Name,
		BucketNamespace: *bucket.Namespace,
		FilePath:        *summary.Name,
	}
}
