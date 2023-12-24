package orchestration

import "github.com/jrolstad/log-analytics-platform/internal/logging"

func ProcessFileInBucket(input interface{}) error {
	logging.LogEvent("Processing file", "data", input)
	return nil
}
