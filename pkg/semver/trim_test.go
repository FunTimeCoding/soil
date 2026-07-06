package semver

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestTrim(t *testing.T) {
	assert.String(t, "0.0.0", Trim("v0.0.0"))
}
