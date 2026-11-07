package option

import (
	"github.com/funtimecoding/soil/pkg/directory/face"
	"github.com/funtimecoding/soil/pkg/web/option"
)

type Directory struct {
	option.Web
	Host   string
	Source func() face.DirectoryClient
}
