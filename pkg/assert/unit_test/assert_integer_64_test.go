package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestInteger64(t *testing.T) {
	assert.Integer64(t, 1, 1)
}
