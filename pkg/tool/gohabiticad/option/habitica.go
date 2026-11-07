package option

import (
	"github.com/funtimecoding/soil/pkg/habitica/face"
	"github.com/funtimecoding/soil/pkg/web/option"
)

type Habitica struct {
	option.Web
	Host   string
	Source func() face.HabiticaClient
}
