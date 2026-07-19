package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestDeviate(t *testing.T) {
	assert.Deviate(t, 1, 1.1, 0.1)
	assert.Deviate(t, -1, -1.1, 0.1)
}
