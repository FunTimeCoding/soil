package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/maps"
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"testing"
)

func TestIntegerKeys(t *testing.T) {
	assert.Integers(
		t,
		[]int{0, 1},
		maps.IntegerKeys(
			map[int]string{0: constant.UpperAlfa, 1: constant.UpperBravo},
		),
	)
}
