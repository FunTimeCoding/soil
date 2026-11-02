package command_context

import (
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/face"
)

type Context struct {
	client   face.Client
	terminal *terminal.Terminal
	instance *string
}
