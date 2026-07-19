package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/underscore"
	"testing"
)

func TestToDash(t *testing.T) {
	assert.String(t, "a-b-c", underscore.ToDash("a_b_c"))
}
