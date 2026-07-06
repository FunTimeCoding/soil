package constant

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(t, "bootstrap", Bootstrap)
	assert.String(t, "bootout", Bootout)
	assert.String(t, "codesign", Codesign)
	assert.String(t, "launchctl", Launchctl)
	assert.String(t, "print", Print)
	assert.String(t, "wdutil", Wdutil)
}
