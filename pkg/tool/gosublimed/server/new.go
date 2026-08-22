package server

import "github.com/funtimecoding/soil/pkg/tool/gosublimed/face"

func New(c face.SublimeSource) *Server {
	return &Server{client: c}
}
