package maps

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"testing"
)

func TestStringKeys(t *testing.T) {
	assert.Strings(
		t,
		[]string{"Alfa", "Bravo"},
		StringKeys(map[string]int{upper.Alfa: 0, upper.Bravo: 1}),
	)
}
