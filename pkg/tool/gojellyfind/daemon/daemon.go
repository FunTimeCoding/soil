package daemon

import (
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/face"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/service"
)

type Daemon struct {
	source  face.JellyfinSource
	service *service.Service
}
