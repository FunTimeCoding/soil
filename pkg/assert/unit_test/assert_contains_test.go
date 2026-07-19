package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestContains(t *testing.T) {
	assert.Contains(t, []string{}, []string{"a", "b"})
	assert.Contains(t, []string{"a"}, []string{"a", "b"})
	assert.Contains(t, []string{"a", "b"}, []string{"a", "b"})
}
