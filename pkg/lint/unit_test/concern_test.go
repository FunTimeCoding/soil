package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"testing"
)

func TestNewLine(t *testing.T) {
	c := concern.NewLine(
		constant.UpperAlfa,
		constant.UpperBravo,
		constant.UpperCharlie,
		1,
		constant.UpperCharlie,
		false,
	)
	assert.NotNil(t, c)
	assert.String(t, "line", c.Type)
	assert.Integer(t, 1, c.Line)
}

func TestNewFile(t *testing.T) {
	c := concern.NewFile(
		constant.UpperAlfa,
		constant.UpperBravo,
		constant.UpperCharlie,
		false,
	)
	assert.NotNil(t, c)
	assert.String(t, "file", c.Type)
	assert.Integer(t, 0, c.Line)
}

func TestNewPackage(t *testing.T) {
	c := concern.NewPackage(
		constant.UpperAlfa,
		constant.UpperBravo,
		constant.UpperCharlie,
	)
	assert.NotNil(t, c)
	assert.String(t, "package", c.Type)
	assert.Boolean(t, false, c.Fixed)
}

func TestNewDelegatesToNewLine(t *testing.T) {
	c := concern.New(
		constant.UpperAlfa,
		constant.UpperBravo,
		constant.UpperCharlie,
		1,
		constant.UpperCharlie,
		false,
	)
	assert.String(t, "line", c.Type)
}
