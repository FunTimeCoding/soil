package server

import (
	"github.com/funtimecoding/soil/pkg/face"
	sublime "github.com/funtimecoding/soil/pkg/tool/gosublimed/face"
)

type Server struct {
	service  sublime.Sublime
	reporter face.Reporter
}
