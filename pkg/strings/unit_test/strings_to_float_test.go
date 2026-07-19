package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings"
	"testing"
)

func TestToFloat(t *testing.T) {
	assert.Float(t, 0.5, strings.ToFloat("0.5", 0))
	assert.Float(t, 0.5, strings.ToFloat(" 0.5", 0))
}
