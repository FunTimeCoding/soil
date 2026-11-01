package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/maps"
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"testing"
)

func TestEncode(t *testing.T) {
	assert.Bytes(
		t,
		[]byte(`{"a":"b"}`),
		maps.Encode(map[string]string{"a": "b"}),
	)
}

func TestIntegerKeys(t *testing.T) {
	assert.Integers(
		t,
		[]int{0, 1},
		maps.IntegerKeys(
			map[int]string{0: constant.UpperAlfa, 1: constant.UpperBravo},
		),
	)
}

func TestStringKeys(t *testing.T) {
	assert.Strings(
		t,
		[]string{"Alfa", "Bravo"},
		maps.StringKeys(
			map[string]int{constant.UpperAlfa: 0, constant.UpperBravo: 1},
		),
	)
}

func TestJoin(t *testing.T) {
	assert.Strings(
		t,
		[]string{"Alfa", "Bravo"},
		maps.StringKeys(
			maps.Join(
				map[string]int{constant.UpperAlfa: 0},
				map[string]int{constant.UpperBravo: 1},
			),
		),
	)
}

func TestJoinLaterPartWins(t *testing.T) {
	assert.MapValue(
		t,
		1,
		maps.Join(
			map[string]any{constant.UpperAlfa: 0},
			map[string]any{constant.UpperAlfa: 1},
		),
		"Alfa",
	)
}

func TestToMarkup(t *testing.T) {
	assert.String(
		t,
		"  a: dGVzdA==\n",
		maps.ToMarkup(map[string]string{"a": "test"}),
	)
}
