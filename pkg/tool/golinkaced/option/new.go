package option

import (
	"github.com/funtimecoding/soil/pkg/linkace"
	"github.com/funtimecoding/soil/pkg/linkace/face"
)

func New() *LinkAce {
	return &LinkAce{
		Source: func() face.Client {
			return linkace.NewEnvironment()
		},
	}
}
