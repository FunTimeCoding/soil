package gomemory

import (
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/face"
)

type Context struct {
	Client   face.Client
	Terminal *terminal.Terminal
}
