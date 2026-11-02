package server

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/diff"
)

type Server struct {
	diff     *diff.Service
	reporter face.Reporter
}
