package option

import (
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/client"
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/face"
	authorization "github.com/funtimecoding/soil/pkg/web/authorization/client"
	"github.com/funtimecoding/soil/pkg/web/option"
)

func New() *Raid {
	return &Raid{
		Web:    option.Web{Authorization: authorization.NewEnvironment},
		Parser: func() face.Client { return client.NewEnvironment() },
	}
}
