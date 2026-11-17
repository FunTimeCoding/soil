package service

import "github.com/funtimecoding/soil/pkg/tool/gosproutd/store/session_statistic"

func (s *Service) SessionStatistics() []*session_statistic.Statistic {
	return s.store.SessionStatistics()
}
