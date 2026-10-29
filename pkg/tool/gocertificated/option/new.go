package option

import (
	"github.com/funtimecoding/soil/pkg/gitlab"
	"github.com/funtimecoding/soil/pkg/gitlab/face"
)

func New() *Certificate {
	return &Certificate{
		Forge: func() face.Forge { return gitlab.NewEnvironment() },
	}
}
