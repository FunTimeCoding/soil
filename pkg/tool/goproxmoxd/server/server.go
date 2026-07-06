package server

import (
	"github.com/funtimecoding/soil/pkg/face"
	proxFace "github.com/funtimecoding/soil/pkg/tool/goproxmoxd/face"
)

type Server struct {
	service  proxFace.Service
	reporter face.Reporter
}
