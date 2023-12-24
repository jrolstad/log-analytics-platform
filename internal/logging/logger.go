package logging

import (
	"go.uber.org/zap"
	"sync"
)

func LogEvent(message string, keysAndValues ...interface{}) {
	logger := getLogger()
	properties := appendLogType(keysAndValues, "event")
	logger.Infow(message, properties...)
}

func LogDependency(message string, keysAndValues ...interface{}) {
	logger := getLogger()
	properties := appendLogType(keysAndValues, "dependency")
	logger.Infow(message, properties...)
}

func LogTrace(message string, keysAndValues ...interface{}) {
	logger := getLogger()
	properties := appendLogType(keysAndValues, "trace")
	logger.Infow(message, properties...)
}

var logger *zap.SugaredLogger
var locker = &sync.Mutex{}

func getLogger() *zap.SugaredLogger {
	if logger == nil {
		locker.Lock()

		prodLogger, _ := zap.NewProduction()
		logger = prodLogger.Sugar()

		locker.Unlock()
	}

	return logger
}

func appendLogType(toMap []interface{}, logType string) []interface{} {
	var result []interface{} = nil
	if toMap == nil {
		result = make([]interface{}, 0)
	} else {
		result = toMap
	}

	result = append(result, "logType", logType)

	return result
}
