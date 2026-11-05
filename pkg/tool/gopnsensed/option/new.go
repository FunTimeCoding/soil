package option

import (
	"github.com/funtimecoding/soil/pkg/opnsense"
	"github.com/funtimecoding/soil/pkg/opnsense/face"
)

func New() *Opnsense {
	return &Opnsense{
		Source: func() face.OpnsenseClient { return opnsense.NewEnvironment() },
	}
}
