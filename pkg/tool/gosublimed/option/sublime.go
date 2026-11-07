package option

import (
	"github.com/funtimecoding/soil/pkg/sublime/face"
	"github.com/funtimecoding/soil/pkg/web/option"
)

type Sublime struct {
	option.Web
	Source func() face.SublimeClient
}
