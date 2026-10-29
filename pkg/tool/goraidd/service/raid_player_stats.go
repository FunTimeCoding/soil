package service

import "github.com/funtimecoding/soil/pkg/tool/goraidd/store/record"

func (s *Service) RaidPlayerStats(raidIdentifier int) []record.RaidPlayer {
	return s.store.RaidPlayerStats(raidIdentifier)
}
