package option

import (
	"github.com/funtimecoding/soil/pkg/habitica"
	"github.com/funtimecoding/soil/pkg/habitica/face"
)

func New() *Habitica {
	return &Habitica{
		Source: func() face.Client {
			return habitica.NewEnvironment()
		},
	}
}
