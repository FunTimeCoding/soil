package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/linux/constant"
	"testing"
)

func TestSystemdConstant(t *testing.T) {
	assert.String(t, "systemctl", constant.SystemdCommand)
}
