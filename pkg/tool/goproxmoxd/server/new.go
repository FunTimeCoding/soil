package server

import (
	"github.com/funtimecoding/soil/pkg/face"
	proxFace "github.com/funtimecoding/soil/pkg/tool/goproxmoxd/face"
)

func New(
	v proxFace.Service,
	r face.Reporter,
) *Server {
	return &Server{service: v, reporter: r}
}
