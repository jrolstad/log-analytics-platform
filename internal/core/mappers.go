package core

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

func MapToJson(toMap interface{}) string {
	result, err := json.Marshal(toMap)
	if err != nil {
		return "{}"
	}

	return string(result)
}

func MapUniqueIdentifier(values ...string) string {
	// TODO: Remove now.string
	values = append(values, time.Now().String())
	resultingValue := strings.Join(values, "|")
	hashedValue := sha256.Sum256([]byte(resultingValue))

	return fmt.Sprintf("%x", hashedValue)

}
