package option

import (
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/client"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/face"
)

func New() *Source {
	return &Source{
		Memory: func() face.Client { return client.NewEnvironment() },
	}
}
