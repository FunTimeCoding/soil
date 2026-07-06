package strings

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"testing"
)

func TestToLower(t *testing.T) {
	assert.Strings(
		t,
		[]string{"alfa", "bravo"},
		ToLower([]string{upper.Alfa, upper.Bravo}),
	)
}
