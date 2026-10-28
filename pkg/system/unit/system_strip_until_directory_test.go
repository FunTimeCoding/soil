package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"testing"
)

func TestStripUntilDirectory(t *testing.T) {
	assert.String(
		t,
		"/test",
		system.StripUntilDirectory(
			"/some/long/path/test",
			constant.FixtureDirectory,
		),
	)
}
