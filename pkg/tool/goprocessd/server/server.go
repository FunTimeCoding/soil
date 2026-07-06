package server

import (
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/environment"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/process"
)

type Server struct {
	processes    []*process.Process
	maxNameWidth int
	environment  *environment.Environment
	procfilePath string
	envrcPath    string
	socketPath   string
}
