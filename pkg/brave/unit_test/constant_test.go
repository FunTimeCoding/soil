package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/brave/constant"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(
		t,
		"/Applications/Brave Browser.app/Contents/MacOS/Brave Browser",
		constant.BravePath,
	)
}
