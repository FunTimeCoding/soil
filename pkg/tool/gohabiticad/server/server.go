package server

import (
	library "github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/face"
)

type Server struct {
	habitica face.HabiticaSource
	reporter library.Reporter
}
