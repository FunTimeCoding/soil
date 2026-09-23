package service

import "github.com/funtimecoding/soil/pkg/tool/gosproutd/store/tally"

func (s *Service) Tallies() []*tally.Tally {
	return s.store.Tallies()
}
