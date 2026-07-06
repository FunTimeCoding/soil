package service

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/store/status"

func (s *Service) MustStatus() *status.Status {
	return s.store.MustStatus()
}
