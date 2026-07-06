package service

import "github.com/funtimecoding/soil/pkg/tool/gosproutd/store/seed"

func (s *Service) Seeds() []seed.Seed {
	return s.store.Seeds()
}
