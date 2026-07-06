package strings

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"testing"
)

func TestAnyToSlice(t *testing.T) {
	assert.Strings(t, []string{"Alfa"}, AnyToSlice(upper.Alfa))
	assert.Strings(
		t,
		[]string{"Alfa", "Bravo"},
		AnyToSlice([]string{upper.Alfa, upper.Bravo}),
	)
}
