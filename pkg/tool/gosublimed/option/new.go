package option

import (
	"github.com/funtimecoding/soil/pkg/sublime"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/face"
)

func New() *Sublime {
	return &Sublime{
		Source: func() face.SublimeSource { return sublime.NewEnvironment() },
	}
}
