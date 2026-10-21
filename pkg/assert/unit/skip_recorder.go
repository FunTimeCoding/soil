package unit

import "testing"

type skipRecorder struct {
	testing.TB
	skipped bool
}
