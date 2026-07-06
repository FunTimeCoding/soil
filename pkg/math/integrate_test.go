package math

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"gonum.org/v1/gonum/integrate"
	"testing"
)

func TestIntegrate(t *testing.T) {
	assert.Round(
		t,
		1,
		integrate.Trapezoidal([]float64{1, 2}, []float64{1, 1}),
		1,
	)
}
