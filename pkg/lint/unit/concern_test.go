package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"testing"
)

func TestNewLine(t *testing.T) {
	c := concern.NewLine(
		stringsConstant.UpperAlfa,
		stringsConstant.UpperBravo,
		stringsConstant.UpperCharlie,
		1,
		stringsConstant.UpperCharlie,
		false,
	)
	assert.NotNil(t, c)
	assert.String(t, "line", c.Type)
	assert.Integer(t, 1, c.Line)
}

func TestNewFile(t *testing.T) {
	c := concern.NewFile(
		stringsConstant.UpperAlfa,
		stringsConstant.UpperBravo,
		stringsConstant.UpperCharlie,
		false,
	)
	assert.NotNil(t, c)
	assert.String(t, "file", c.Type)
	assert.Integer(t, 0, c.Line)
}

func TestNewPackage(t *testing.T) {
	c := concern.NewPackage(
		stringsConstant.UpperAlfa,
		stringsConstant.UpperBravo,
		stringsConstant.UpperCharlie,
	)
	assert.NotNil(t, c)
	assert.String(t, "package", c.Type)
	assert.Boolean(t, false, c.Fixed)
}

func TestNewDelegatesToNewLine(t *testing.T) {
	c := concern.New(
		stringsConstant.UpperAlfa,
		stringsConstant.UpperBravo,
		stringsConstant.UpperCharlie,
		1,
		stringsConstant.UpperCharlie,
		false,
	)
	assert.String(t, "line", c.Type)
}
