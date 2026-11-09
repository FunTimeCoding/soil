package option

import (
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/client"
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/face"
	authorization "github.com/funtimecoding/soil/pkg/web/authorization/client"
	"github.com/funtimecoding/soil/pkg/web/option"
	"time"
)

func New() *Raid {
	return &Raid{
		Web:    option.Web{Authorization: authorization.NewEnvironment},
		Clock:  time.Now,
		Parser: func() face.Client { return client.NewEnvironment() },
	}
}
