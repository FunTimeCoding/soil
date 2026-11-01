package service

import "github.com/funtimecoding/soil/pkg/tool/goagentd/types/run_result"

func (s *Service) Status() *run_result.Result {
	return s.runner.Status()
}
