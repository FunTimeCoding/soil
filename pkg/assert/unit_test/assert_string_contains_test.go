package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestStringContains(t *testing.T) {
	assert.StringContains(t, "friend", "hello friend")
}
