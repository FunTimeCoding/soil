package option

import (
	"github.com/funtimecoding/soil/pkg/github"
	"github.com/funtimecoding/soil/pkg/tool/gogithubd/face"
)

func New() *Exporter {
	return &Exporter{
		Source: func() face.RunSource {
			return github.NewEnvironment()
		},
	}
}
