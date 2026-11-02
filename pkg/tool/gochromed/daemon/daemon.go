package daemon

import (
	"github.com/funtimecoding/soil/pkg/tool/gochromed/face"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/service"
)

type Daemon struct {
	source  face.ChromiumSource
	service *service.Service
}
