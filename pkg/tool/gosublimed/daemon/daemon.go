package daemon

import (
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/face"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/service"
)

type Daemon struct {
	source  face.SublimeSource
	service *service.Service
}
