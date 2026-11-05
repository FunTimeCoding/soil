package command_context

import (
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/face"
)

func New(
	t *terminal.Terminal,
	client face.Client,
	longClient face.Client,
) *Context {
	return &Context{terminal: t, client: client, longClient: longClient}
}
