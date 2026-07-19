package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"testing"
)

func TestAnyToSlice(t *testing.T) {
	assert.Strings(t, []string{"Alfa"}, strings.AnyToSlice(upper.Alfa))
	assert.Strings(
		t,
		[]string{"Alfa", "Bravo"},
		strings.AnyToSlice([]string{upper.Alfa, upper.Bravo}),
	)
}
