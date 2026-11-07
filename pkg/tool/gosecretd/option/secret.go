package option

import (
	"github.com/funtimecoding/soil/pkg/openbao/face"
	"github.com/funtimecoding/soil/pkg/web/option"
)

type Secret struct {
	option.Web
	Metric option.Web
	Source func() face.OpenbaoClient
}
