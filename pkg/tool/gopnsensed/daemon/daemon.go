package daemon

import (
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/face"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/service"
)

type Daemon struct {
	source  face.OpnsenseSource
	service *service.Service
}
