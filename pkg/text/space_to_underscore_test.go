package text

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestSpaceToUnderscore(t *testing.T) {
	assert.String(t, "a_b_c", SpaceToUnderscore("a b c"))
}
