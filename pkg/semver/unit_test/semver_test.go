package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/semver"
	"testing"
)

func TestTrim(t *testing.T) {
	assert.String(t, "0.0.0", semver.Trim("v0.0.0"))
}
