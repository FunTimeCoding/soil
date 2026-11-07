package option

import (
	"github.com/funtimecoding/soil/pkg/directory"
	"github.com/funtimecoding/soil/pkg/directory/face"
	"github.com/funtimecoding/soil/pkg/web/authorization/client"
	"github.com/funtimecoding/soil/pkg/web/option"
)

func New() *Directory {
	return &Directory{
		Web: option.Web{Authorization: client.NewEnvironment},
		Source: func() face.DirectoryClient {
			return directory.NewEnvironment()
		},
	}
}
