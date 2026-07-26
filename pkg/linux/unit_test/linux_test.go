package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/linux/constant"
	"testing"
)

func TestLinuxConstant(t *testing.T) {
	assert.String(t, "jc", constant.Jc)
}
