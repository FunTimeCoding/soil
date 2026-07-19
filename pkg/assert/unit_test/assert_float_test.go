package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestFloat(t *testing.T) {
	assert.Float(t, 1.1, 1.1)
	assert.Float(t, 1.15, 1.15)
}
