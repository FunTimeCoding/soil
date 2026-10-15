package server

import (
	"github.com/funtimecoding/soil/pkg/face"
	provisionFace "github.com/funtimecoding/soil/pkg/provision/face"
)

func New(
	n provisionFace.Runner,
	r face.Reporter,
) *Server {
	return &Server{runner: n, reporter: r}
}
