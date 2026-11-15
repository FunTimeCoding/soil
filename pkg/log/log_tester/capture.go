package log_tester

import (
	"bytes"
	"log"
	"testing"
)

func Capture(t *testing.T) *bytes.Buffer {
	t.Helper()
	result := &bytes.Buffer{}
	previous := log.Writer()
	log.SetOutput(result)
	t.Cleanup(func() { log.SetOutput(previous) })

	return result
}
