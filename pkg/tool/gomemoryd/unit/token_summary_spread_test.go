package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/token_summary"
	"testing"
)

func TestSpread(t *testing.T) {
	s := token_summary.NewSpread([]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100})
	assert.Integer(t, 50, s.Median)
	assert.Integer(t, 100, s.Maximum)
	assert.Integer(t, 550, s.Total)
}

func TestSpreadIgnoresInputOrder(t *testing.T) {
	s := token_summary.NewSpread([]int{100, 10, 50, 20, 90})
	assert.Integer(t, 100, s.Maximum)
	assert.Integer(t, 50, s.Median)
}

func TestSpreadDoesNotMutateInput(t *testing.T) {
	values := []int{30, 10, 20}
	token_summary.NewSpread(values)
	assert.Integers(t, []int{30, 10, 20}, values)
}

func TestSpreadEmpty(t *testing.T) {
	s := token_summary.NewSpread(nil)
	assert.Integer(t, 0, s.Maximum)
	assert.Integer(t, 0, s.Total)
}

func TestSpreadSingleValue(t *testing.T) {
	s := token_summary.NewSpread([]int{42})
	assert.Integer(t, 42, s.Median)
	assert.Integer(t, 42, s.Maximum)
	assert.Integer(t, 42, s.Total)
}
