package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/contains"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"testing"
)

func TestSubAny(t *testing.T) {
	assert.True(t, contains.AnySub([]string{"Al"}, []string{upper.Alfa}))
	assert.False(t, contains.AnySub([]string{"Ga"}, []string{upper.Alfa}))
	assert.False(t, contains.AnySub([]string{}, []string{upper.Alfa}))
	assert.False(t, contains.AnySub([]string{}, []string{}))
	assert.False(t, contains.AnySub([]string{"Al"}, []string{}))
	assert.True(
		t,
		contains.AnySub(
			[]string{"Br"},
			[]string{upper.Alfa, upper.Bravo, upper.Charlie},
		),
	)
	assert.False(
		t,
		contains.AnySub(
			[]string{"De"},
			[]string{upper.Alfa, upper.Bravo, upper.Charlie},
		),
	)
}
