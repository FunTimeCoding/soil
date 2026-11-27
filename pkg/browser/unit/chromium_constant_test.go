package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/browser/constant"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(t, "CHROMIUM_HOST", constant.ChromiumHostEnvironment)
}
