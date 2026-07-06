package math

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestSeriesSum(t *testing.T) {
	assertSeriesSum(t, 1, 1)
	assertSeriesSum(t, 9, 45)
}

func assertSeriesSum(
	t *testing.T,
	n int,
	expect int,
) {
	t.Helper()
	assert.Integer(t, expect, SeriesSum(n))
}
