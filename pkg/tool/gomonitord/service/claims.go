package service

import "github.com/funtimecoding/soil/pkg/tool/gomonitord/store/claim"

func (s *Service) Claims() ([]*claim.Claim, error) {
	return s.store.Claims()
}
