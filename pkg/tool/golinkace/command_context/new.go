package command_context

import (
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/web/stall_client"
)

func New(t *terminal.Terminal) *Context {
	return &Context{terminal: t, pool: stall_client.New().Build()}
}
