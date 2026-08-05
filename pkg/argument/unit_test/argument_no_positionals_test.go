package unit_test

import (
	"github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestNoPositionalsAccepts(t *testing.T) {
	a := testInstance(t)
	assert.Nil(t, a.ParseArguments(nil))
	a.NoPositionals("hint")
}

func TestNoPositionalsAcceptsFlags(t *testing.T) {
	a := testInstance(t)
	a.String(constant.File, "Procfile", constant.Path)
	assert.Nil(t, a.ParseArguments([]string{"--file", "Other"}))
	a.NoPositionals("hint")
}
