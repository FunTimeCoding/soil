package service

import "github.com/funtimecoding/soil/pkg/tool/goflightd/store/mark"

func (s *Service) Marks(limit int) ([]mark.Mark, error) {
	return s.store.RecentMarks(limit)
}
