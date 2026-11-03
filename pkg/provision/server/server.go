package server

import (
	soilFace "github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/provision/face"
)

type Server struct {
	runner   face.Runner
	reporter soilFace.Reporter
}
