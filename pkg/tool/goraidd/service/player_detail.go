package service

import "github.com/funtimecoding/soil/pkg/tool/goraidd/store/record"

func (s *Service) PlayerDetail(account string) []record.PlayerRaid {
	return s.store.PlayerDetail(account)
}
