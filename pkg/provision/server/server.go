package server

import (
	"github.com/funtimecoding/soil/pkg/face"
	provisionFace "github.com/funtimecoding/soil/pkg/provision/face"
)

type Server struct {
	runner   provisionFace.Runner
	reporter face.Reporter
}
