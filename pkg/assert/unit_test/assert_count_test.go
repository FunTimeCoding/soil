package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestCount(t *testing.T) {
	assert.Count(t, 0, []string{})
	assert.Count(t, 1, []string{"1"})
	assert.Count(t, 2, []string{"1", "2"})
}
