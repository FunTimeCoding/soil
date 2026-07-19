package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings"
	"testing"
)

func TestMustToFloat(t *testing.T) {
	assert.Float(t, 1, strings.MustToFloat("1.0"))
	assert.Float(t, 1, strings.MustToFloat(" 1.0"))
}
