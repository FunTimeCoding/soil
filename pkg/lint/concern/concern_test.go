package concern

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"testing"
)

func TestNewLine(t *testing.T) {
	c := NewLine(
		upper.Alfa,
		upper.Bravo,
		upper.Charlie,
		1,
		upper.Charlie,
		false,
	)
	assert.NotNil(t, c)
	assert.String(t, Line, c.Type)
	assert.Integer(t, 1, c.Line)
}

func TestNewFile(t *testing.T) {
	c := NewFile(
		upper.Alfa,
		upper.Bravo,
		upper.Charlie,
		false,
	)
	assert.NotNil(t, c)
	assert.String(t, File, c.Type)
	assert.Integer(t, 0, c.Line)
}

func TestNewPackage(t *testing.T) {
	c := NewPackage(
		upper.Alfa,
		upper.Bravo,
		upper.Charlie,
	)
	assert.NotNil(t, c)
	assert.String(t, Package, c.Type)
	assert.Boolean(t, false, c.Fixed)
}

func TestNewDelegatesToNewLine(t *testing.T) {
	c := New(
		upper.Alfa,
		upper.Bravo,
		upper.Charlie,
		1,
		upper.Charlie,
		false,
	)
	assert.String(t, Line, c.Type)
}
