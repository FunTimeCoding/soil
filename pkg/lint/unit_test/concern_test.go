package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/constant"
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
	assert.String(t, constant.ConcernLine, c.Type)
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
	assert.String(t, constant.ConcernFile, c.Type)
	assert.Integer(t, 0, c.Line)
}

func TestNewPackage(t *testing.T) {
	c := concern.NewPackage(
		stringsConstant.UpperAlfa,
		stringsConstant.UpperBravo,
		stringsConstant.UpperCharlie,
	)
	assert.NotNil(t, c)
	assert.String(t, constant.ConcernPackage, c.Type)
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
	assert.String(t, constant.ConcernLine, c.Type)
}
