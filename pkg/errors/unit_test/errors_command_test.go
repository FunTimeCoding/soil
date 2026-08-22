package unit_test

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/command"
	"github.com/funtimecoding/soil/pkg/errors/constant"
	"github.com/funtimecoding/soil/pkg/face"
	"testing"
)

func TestCommandErrorCarriesTheStory(t *testing.T) {
	wrapped := errors.New("exit status 128")
	e := command.New(
		"git reset --hard origin/main",
		"",
		"fatal: index.lock exists",
		wrapped,
	)
	assert.String(t, "git reset --hard origin/main: exit status 128", e.Error())
	assert.True(t, errors.Is(e, wrapped))
	assert.True(t, command.Is(e))
	var provider face.ContextProvider = e
	key, context := provider.ErrorContext()
	assert.String(t, "process", key)
	assert.String(
		t,
		"git reset --hard origin/main",
		context[constant.Command].(string),
	)
	assert.String(
		t,
		"fatal: index.lock exists",
		context[constant.Stderr].(string),
	)
}
