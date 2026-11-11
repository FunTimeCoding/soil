package option

import (
	"github.com/funtimecoding/soil/pkg/backrest"
	"github.com/funtimecoding/soil/pkg/backrest/face"
)

func New() *Exporter {
	return &Exporter{
		Source: func() face.BackrestClient {
			return backrest.NewEnvironment()
		},
	}
}
