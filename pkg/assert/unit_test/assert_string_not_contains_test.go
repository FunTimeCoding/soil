package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestStringNotContains(t *testing.T) {
	assert.StringNotContains(t, "enemy", "hello friend")
}
