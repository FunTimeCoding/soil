package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/forge/constant"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(t, "FORGE_AUTO_CLEANUP", constant.AutoCleanupEnvironment)
}
