package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings"
	"testing"
)

func TestToInteger(t *testing.T) {
	assert.Integer(t, 5, strings.ToInteger(strings.Five, 0))
	assert.Integer(t, 5, strings.ToInteger(" 5", 0))
}
