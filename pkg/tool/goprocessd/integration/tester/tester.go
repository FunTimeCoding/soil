package tester

import "github.com/funtimecoding/soil/pkg/tool/goprocessd/supervisor"

type Tester struct {
	Server       *supervisor.Supervisor
	Directory    string
	SocketPath   string
	ProcfilePath string
	EnvrcPath    string
}
