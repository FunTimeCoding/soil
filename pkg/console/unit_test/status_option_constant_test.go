package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/console/constant"
	"testing"
)

func TestStatusOptionConstant(t *testing.T) {
	assert.NotNil(t, constant.ColorFormat)
	assert.NotNil(t, constant.ExtendedColorFormat)
}
