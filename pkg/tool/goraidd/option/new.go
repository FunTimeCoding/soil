package option

import (
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/client"
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/face"
)

func New() *Raid {
	return &Raid{Parser: func() face.Client { return client.NewEnvironment() }}
}
