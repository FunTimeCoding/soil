package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/contains"
	"testing"
)

func TestAll(t *testing.T) {
	assert.True(
		t,
		contains.All(
			[]string{constant.UpperAlfa},
			[]string{constant.UpperAlfa},
		),
	)
	assert.False(
		t,
		contains.All(
			[]string{constant.UpperAlfa},
			[]string{constant.UpperBravo},
		),
	)
	assert.True(
		t,
		contains.All(
			[]string{
				constant.UpperAlfa,
				constant.UpperBravo,
				constant.UpperCharlie,
			},
			[]string{constant.UpperAlfa, constant.UpperBravo},
		),
	)
	assert.False(
		t,
		contains.All(
			[]string{
				constant.UpperAlfa,
				constant.UpperBravo,
				constant.UpperCharlie,
			},
			[]string{constant.UpperAlfa, constant.UpperDelta},
		),
	)
}

func TestAny(t *testing.T) {
	assert.True(
		t,
		contains.Any(
			[]string{constant.UpperAlfa},
			[]string{constant.UpperAlfa},
		),
	)
	assert.False(
		t,
		contains.Any(
			[]string{constant.UpperAlfa},
			[]string{constant.UpperBravo},
		),
	)
	assert.True(
		t,
		contains.Any(
			[]string{
				constant.UpperAlfa,
				constant.UpperBravo,
				constant.UpperCharlie,
			},
			[]string{constant.UpperAlfa},
		),
	)
	assert.False(
		t,
		contains.Any(
			[]string{
				constant.UpperAlfa,
				constant.UpperBravo,
				constant.UpperCharlie,
			},
			[]string{constant.UpperDelta},
		),
	)
}

func TestSubAny(t *testing.T) {
	assert.True(
		t,
		contains.AnySub([]string{"Al"}, []string{constant.UpperAlfa}),
	)
	assert.False(
		t,
		contains.AnySub([]string{"Ga"}, []string{constant.UpperAlfa}),
	)
	assert.False(t, contains.AnySub([]string{}, []string{constant.UpperAlfa}))
	assert.False(t, contains.AnySub([]string{}, []string{}))
	assert.False(t, contains.AnySub([]string{"Al"}, []string{}))
	assert.True(
		t,
		contains.AnySub(
			[]string{"Br"},
			[]string{
				constant.UpperAlfa,
				constant.UpperBravo,
				constant.UpperCharlie,
			},
		),
	)
	assert.False(
		t,
		contains.AnySub(
			[]string{"De"},
			[]string{
				constant.UpperAlfa,
				constant.UpperBravo,
				constant.UpperCharlie,
			},
		),
	)
}
