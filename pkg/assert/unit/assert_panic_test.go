package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestPanic(t *testing.T) {
	assert.Panic(t, func() { panic("alfa") })
}

func TestPanicRejectsReturning(t *testing.T) {
	inner := &testing.T{}
	assert.Panic(inner, func() {})
	assert.True(t, inner.Failed())
}

func TestPanicContains(t *testing.T) {
	assert.PanicContains(t, "alfa", func() { panic("alfa bravo") })
	assert.PanicContains(t, "bravo", func() { panic(errors.New("alfa bravo")) })
}

func TestPanicContainsRejectsReturning(t *testing.T) {
	inner := &testing.T{}
	assert.PanicContains(inner, "alfa", func() {})
	assert.True(t, inner.Failed())
}

func TestPanicContainsRejectsOtherMessage(t *testing.T) {
	inner := &testing.T{}
	assert.PanicContains(inner, "charlie", func() { panic("alfa bravo") })
	assert.True(t, inner.Failed())
}
