package server

import (
	library "github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/face"
)

func New(
	c face.HabiticaSource,
	r library.Reporter,
) *Server {
	return &Server{habitica: c, reporter: r}
}
