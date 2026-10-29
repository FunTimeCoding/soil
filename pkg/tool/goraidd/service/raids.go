package service

import "github.com/funtimecoding/soil/pkg/tool/goraidd/store/record"

func (s *Service) Raids() []record.Raid {
	return s.store.Raids()
}
