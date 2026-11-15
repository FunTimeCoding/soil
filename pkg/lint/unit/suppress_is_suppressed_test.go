package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/unit/suppress_tester"
	"testing"
)

func TestDirectiveOnFirstLineSuppressesCall(t *testing.T) {
	o := suppress_tester.New(t, constant.SuppressCallSample)
	assert.True(t, o.Suppressed(o.Call(7), "forbidden_call"))
}

func TestDirectiveOnLastLineSuppressesCall(t *testing.T) {
	o := suppress_tester.New(t, constant.SuppressCallSample)
	assert.True(t, o.Suppressed(o.Call(4), "forbidden_call"))
}

func TestCallUnsuppressedByDirectiveOnArgumentLine(t *testing.T) {
	o := suppress_tester.New(t, constant.SuppressCallSample)
	assert.False(t, o.Suppressed(o.Call(8), "forbidden_call"))
	assert.True(t, o.Suppressed(o.Call(9), "forbidden_call"))
}

func TestCallUnsuppressedByDirectiveInsideClosure(t *testing.T) {
	o := suppress_tester.New(t, constant.SuppressCallSample)
	assert.False(t, o.Suppressed(o.Call(12), "forbidden_call"))
	assert.True(t, o.Suppressed(o.Call(13), "forbidden_call"))
}

func TestDirectiveBelowOpeningLineSuppressesLiteral(t *testing.T) {
	o := suppress_tester.New(t, constant.SuppressLiteralSample)
	assert.True(t, o.Suppressed(o.Literal(4), "restricted_literal"))
	assert.True(t, o.Suppressed(o.Literal(6), "restricted_literal"))
}

func TestDirectiveOnLastLineSuppressesLiteral(t *testing.T) {
	o := suppress_tester.New(t, constant.SuppressLiteralSample)
	assert.True(t, o.Suppressed(o.Literal(8), "restricted_literal"))
	assert.True(t, o.Suppressed(o.Literal(9), "restricted_literal"))
}

func TestLiteralUnsuppressedByInnerLiteralDirective(t *testing.T) {
	o := suppress_tester.New(t, constant.SuppressLiteralSample)
	assert.False(t, o.Suppressed(o.Literal(13), "restricted_literal"))
	assert.True(t, o.Suppressed(o.Literal(14), "restricted_literal"))
}
