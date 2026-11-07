package option

import (
	"github.com/funtimecoding/soil/pkg/opnsense/face"
	"github.com/funtimecoding/soil/pkg/web/option"
)

type Opnsense struct {
	option.Web
	Source func() face.OpnsenseClient
}
