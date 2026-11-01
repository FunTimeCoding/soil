package service

import "github.com/funtimecoding/soil/pkg/tool/goclauded/sweep"

func (s *Service) Sweep() *sweep.Result {
	return sweep.Run(s.harbor)
}
