package server

import (
	"github.com/funtimecoding/soil/pkg/face"
	parser "github.com/funtimecoding/soil/pkg/tool/goraidparsed/face"
)

type Server struct {
	service  parser.Generator
	reporter face.Reporter
}
