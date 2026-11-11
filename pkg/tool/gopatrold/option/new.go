package option

import (
	"github.com/funtimecoding/soil/pkg/patrol"
	"github.com/funtimecoding/soil/pkg/patrol/face"
)

func New() *Patrol {
	return &Patrol{
		Source: func(host string) face.PatrolClient {
			return patrol.NewEnvironment(host)
		},
	}
}
