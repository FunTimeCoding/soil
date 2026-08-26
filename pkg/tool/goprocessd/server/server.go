package server

import (
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/environment"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/process"
	"sync"
)

type Server struct {
	processes    []*process.Process
	maxNameWidth int
	environment  *environment.Environment
	procfilePath string
	envrcPath    string
	socketPath   string
	running      int
	countMutex   sync.Mutex
	processMutex sync.RWMutex
	commandMutex sync.Mutex
	allDone      chan struct{}
}
