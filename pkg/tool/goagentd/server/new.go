package server

import "github.com/funtimecoding/soil/pkg/tool/goagentd/runner"

func New(runner *runner.Runner) *Server {
	return &Server{runner: runner}
}
