package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestNotEmpty(t *testing.T) {
	assert.NotEmpty(t, []string{"one"})
	assert.NotEmpty(t, map[string]int{"a": 1})
}
