package service

import "github.com/funtimecoding/soil/pkg/tool/goprocessd/status"

func (s *Service) Statuses() []*status.Status {
	return s.supervisor.Statuses()
}
