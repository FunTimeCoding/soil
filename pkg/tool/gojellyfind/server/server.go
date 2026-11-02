package server

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/service"
)

type Server struct {
	service  *service.Service
	reporter face.Reporter
}
