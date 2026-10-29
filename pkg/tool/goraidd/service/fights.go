package service

import "github.com/funtimecoding/soil/pkg/raid/model/fight"

func (s *Service) Fights() []fight.Fight {
	return s.store.Fights()
}
