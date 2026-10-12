package goband

import (
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/generated/client"
)

type Context struct {
	Client   *client.ClientWithResponses
	Terminal *terminal.Terminal
}
