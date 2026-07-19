package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/list"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"testing"
)

func TestJoin(t *testing.T) {
	assert.Any(
		t,
		[]string{"Alfa", "Bravo"},
		list.Join([][]string{{upper.Alfa}, {upper.Bravo}}),
	)
}
