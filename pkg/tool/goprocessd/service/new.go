package service

import "github.com/funtimecoding/soil/pkg/tool/goprocessd/supervisor"

func New(s *supervisor.Supervisor) *Service {
	return &Service{supervisor: s}
}
