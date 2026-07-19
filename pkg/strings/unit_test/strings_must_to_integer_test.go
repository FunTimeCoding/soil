package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings"
	"testing"
)

func TestMustToInteger(t *testing.T) {
	assert.Integer(t, 1, strings.MustToInteger(strings.One))
	assert.Integer(t, 1, strings.MustToInteger(" 1"))
}
