package main

import (
	"context"
	fdk "github.com/fnproject/fdk-go"
	"github.com/jrolstad/log-analytics-platform/internal/logging"
	"io"
)

func main() {
	fdk.Handle(fdk.HandlerFunc(myHandler))
}

func myHandler(ctx context.Context, in io.Reader, out io.Writer) {
	input, err := io.ReadAll(in)
	if err != nil {
		logging.LogError(err)
	}

	logging.LogEvent("Received message", "input", input)
}
