package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/chromium/constant"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(t, "CHROMIUM_HOST", constant.HostEnvironment)
}
