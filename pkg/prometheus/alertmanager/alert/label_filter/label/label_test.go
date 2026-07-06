package label

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"testing"
)

func TestLabel(t *testing.T) {
	assert.True(
		t,
		New(upper.Alfa, upper.Bravo).Match(upper.Alfa, upper.Bravo),
	)
}
