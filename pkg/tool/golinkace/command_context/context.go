package command_context

import (
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/generated/client"
)

type Context struct {
	host     string
	client   *client.ClientWithResponses
	terminal *terminal.Terminal
}
