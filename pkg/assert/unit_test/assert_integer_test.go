package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestInteger(t *testing.T) {
	assert.Integer(t, 1, 1)
}
