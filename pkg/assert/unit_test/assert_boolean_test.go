package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestBoolean(t *testing.T) {
	assert.Boolean(t, true, true)
}
