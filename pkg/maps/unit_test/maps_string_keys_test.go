package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/maps"
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"testing"
)

func TestStringKeys(t *testing.T) {
	assert.Strings(
		t,
		[]string{"Alfa", "Bravo"},
		maps.StringKeys(
			map[string]int{constant.UpperAlfa: 0, constant.UpperBravo: 1},
		),
	)
}
