package math

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestWeight(t *testing.T) {
	assertWeight(t, 1, 1, 1, 1, 1)
	assertWeight(t, 1, 0.5, 1, 1, 0.75)
	assertWeight(t, 1, 0, 1, 1, 0.5)
	assertWeight(t, 255, 0, 1, 1, 127.5)
	assertWeight(t, 1, 0, 2, 1, 0.67)
	assertWeight(t, 1, 0, 1, 2, 0.33)
}

func assertWeight(
	t *testing.T,
	a float64,
	b float64,
	aWeight float64,
	bWeight float64,
	expect float64,
) {
	t.Helper()
	assert.Round(t, expect, Weight(a, b, aWeight, bWeight), 2)
}
