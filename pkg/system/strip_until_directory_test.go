package system

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestStripUntilDirectory(t *testing.T) {
	assert.String(
		t,
		"/test",
		StripUntilDirectory("/some/long/path/test", "test"),
	)
}
