package tester

import "github.com/funtimecoding/soil/pkg/tool/goprocessd/server"

type Tester struct {
	Server       *server.Server
	Directory    string
	SocketPath   string
	ProcfilePath string
	EnvrcPath    string
}
