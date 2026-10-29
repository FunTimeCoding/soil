package server

import (
	"github.com/funtimecoding/soil/pkg/face"
	monitor "github.com/funtimecoding/soil/pkg/tool/gomonitord/face"
)

type Server struct {
	service  monitor.Claims
	reporter face.Reporter
}
