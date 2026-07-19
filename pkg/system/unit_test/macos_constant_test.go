package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system/macos/constant"
	"testing"
)

func TestMacosConstant(t *testing.T) {
	assert.String(t, "bootstrap", constant.Bootstrap)
	assert.String(t, "bootout", constant.Bootout)
	assert.String(t, "codesign", constant.Codesign)
	assert.String(t, "launchctl", constant.Launchctl)
	assert.String(t, "print", constant.Print)
	assert.String(t, "wdutil", constant.Wdutil)
}
