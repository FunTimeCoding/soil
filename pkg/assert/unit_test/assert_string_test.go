package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestString(t *testing.T) {
	assert.String(t, "a", "a")
}
