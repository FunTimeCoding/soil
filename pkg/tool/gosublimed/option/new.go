package option

import (
	"github.com/funtimecoding/soil/pkg/sublime"
	"github.com/funtimecoding/soil/pkg/sublime/face"
)

func New() *Sublime {
	return &Sublime{
		Source: func() face.SublimeClient { return sublime.NewEnvironment() },
	}
}
