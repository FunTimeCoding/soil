package goband

import (
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/face"
)

type Context struct {
	Client   face.Client
	Terminal *terminal.Terminal
}
