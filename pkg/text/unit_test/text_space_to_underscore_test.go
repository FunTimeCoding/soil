package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/text"
	"testing"
)

func TestSpaceToUnderscore(t *testing.T) {
	assert.String(t, "a_b_c", text.SpaceToUnderscore("a b c"))
}
