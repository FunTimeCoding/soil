package service

import "github.com/funtimecoding/soil/pkg/tool/goatlasd/store/result"

func (s *Service) Places() ([]*result.Place, error) {
	return s.store.Places()
}
