package service

import "github.com/funtimecoding/soil/pkg/tool/godashboardd/types/summary"

func (s *Service) Summaries() ([]summary.Summary, error) {
	return s.store.Summaries()
}
