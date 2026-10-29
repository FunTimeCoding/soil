package option

import (
	"github.com/funtimecoding/soil/pkg/openbao"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/face"
)

func New() *Secret {
	return &Secret{
		Source: func() face.SecretSource {
			return openbao.NewEnvironment()
		},
	}
}
