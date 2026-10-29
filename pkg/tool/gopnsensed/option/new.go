package option

import (
	"github.com/funtimecoding/soil/pkg/opnsense"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/face"
)

func New() *Opnsense {
	return &Opnsense{
		Source: func() face.OpnsenseSource { return opnsense.NewEnvironment() },
	}
}
