package service

import "github.com/funtimecoding/soil/pkg/raid/model/fight"

func (s *Service) RaidFights(raidIdentifier int) []fight.Fight {
	return s.store.RaidFights(raidIdentifier)
}
