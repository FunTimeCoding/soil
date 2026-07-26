package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/expression"
	"testing"
)

func TestExpression(t *testing.T) {
	e := expression.New(
		[]string{constant.UpperAlfa},
		[]string{constant.UpperDelta},
	)
	assert.True(t, e.Check([]string{constant.UpperAlfa}))
	assert.True(t, e.Check([]string{constant.UpperAlfa, constant.UpperBravo}))
	assert.False(t, e.Check([]string{constant.UpperAlfa, constant.UpperDelta}))
}
