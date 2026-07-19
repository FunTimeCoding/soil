package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/list"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"testing"
)

type Fixture string

func (f Fixture) String() string {
	return string(f)
}

func TestToStrings(t *testing.T) {
	assert.Strings(
		t,
		[]string{"Alfa", "Bravo"},
		list.ToStrings([]Fixture{upper.Alfa, upper.Bravo}),
	)
}
