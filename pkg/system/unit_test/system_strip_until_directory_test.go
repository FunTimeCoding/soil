package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system"
	"testing"
)

func TestStripUntilDirectory(t *testing.T) {
	assert.String(
		t,
		"/test",
		system.StripUntilDirectory("/some/long/path/test", "test"),
	)
}
