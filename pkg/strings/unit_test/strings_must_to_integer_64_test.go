package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings"
	"testing"
)

func TestMustToInteger64(t *testing.T) {
	assert.Integer64(t, 1, strings.MustToInteger64(strings.One))
	assert.Integer64(t, 1, strings.MustToInteger64(" 1"))
}
