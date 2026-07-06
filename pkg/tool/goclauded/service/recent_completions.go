package service

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/completion"

func (s *Service) RecentCompletions() ([]completion.Completion, error) {
	return s.store.RecentCompletions()
}
