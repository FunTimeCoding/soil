package gosublime

import (
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/face"
)

type Context struct {
	Client   face.Client
	Terminal *terminal.Terminal
}
