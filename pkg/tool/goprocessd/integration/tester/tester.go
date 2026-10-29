package tester

import "github.com/funtimecoding/soil/pkg/tool/goprocessd/client"

type Tester struct {
	Client       *client.Client
	Directory    string
	ProcfilePath string
	EnvrcPath    string
}
