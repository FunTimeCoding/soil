package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/maps"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"testing"
)

func TestStringKeys(t *testing.T) {
	assert.Strings(
		t,
		[]string{"Alfa", "Bravo"},
		maps.StringKeys(map[string]int{upper.Alfa: 0, upper.Bravo: 1}),
	)
}
