package unit

import (
	"bytes"
	"log"
	"os"
	"testing"
)

func captureLog(t *testing.T) *bytes.Buffer {
	t.Helper()
	result := &bytes.Buffer{}
	log.SetOutput(result)
	t.Cleanup(func() { log.SetOutput(os.Stderr) })

	return result
}
