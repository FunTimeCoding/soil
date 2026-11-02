package service

import "github.com/funtimecoding/soil/pkg/tool/goagentd/runner"

func New(r *runner.Runner) *Service {
	return &Service{runner: r}
}
