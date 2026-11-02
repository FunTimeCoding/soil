package daemon

import (
	"github.com/funtimecoding/soil/pkg/tool/goitermd/face"
	"github.com/funtimecoding/soil/pkg/tool/goitermd/service"
)

type Daemon struct {
	source  face.ItermSource
	service *service.Service
}
