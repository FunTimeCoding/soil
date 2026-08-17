package server

import (
	library "github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/face"
)

type Server struct {
	opnsense face.OpnsenseSource
	reporter library.Reporter
}
