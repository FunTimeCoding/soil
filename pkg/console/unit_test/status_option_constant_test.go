package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"testing"
)

func TestStatusOptionConstant(t *testing.T) {
	assert.NotNil(t, option.Color)
	assert.NotNil(t, option.ExtendedColor)
}
