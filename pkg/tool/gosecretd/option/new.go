package option

import (
	"github.com/funtimecoding/soil/pkg/openbao"
	"github.com/funtimecoding/soil/pkg/openbao/face"
)

func New() *Secret {
	return &Secret{
		Source: func() face.OpenbaoClient {
			return openbao.NewEnvironment()
		},
	}
}
