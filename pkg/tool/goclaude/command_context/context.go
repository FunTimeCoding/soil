package command_context

import (
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/face"
)

type Context struct {
	client     face.Client
	longClient face.Client
	terminal   *terminal.Terminal
}
