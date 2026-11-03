package server

import (
	soilFace "github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/provision/face"
)

func New(
	n face.Runner,
	r soilFace.Reporter,
) *Server {
	return &Server{runner: n, reporter: r}
}
