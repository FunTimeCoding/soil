package unit_test

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry"
	rawSentry "github.com/getsentry/sentry-go"
	"testing"
)

func TestCaptureOnError(t *testing.T) {
	sentry.CaptureOnError(rawSentry.CurrentHub(), nil)
}
