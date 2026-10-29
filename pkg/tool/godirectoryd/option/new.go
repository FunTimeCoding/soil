package option

import (
	"github.com/funtimecoding/soil/pkg/directory"
	"github.com/funtimecoding/soil/pkg/directory/face"
)

func New() *Directory {
	return &Directory{
		Source: func() face.Client {
			return directory.NewEnvironment()
		},
	}
}
