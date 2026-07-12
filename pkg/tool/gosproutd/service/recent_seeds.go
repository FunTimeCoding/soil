package service

import "github.com/funtimecoding/soil/pkg/tool/gosproutd/store/seed"

func (s *Service) RecentSeeds() []seed.Seed {
	return s.store.RecentSeeds()
}
