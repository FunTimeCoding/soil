package service

import "github.com/funtimecoding/go-library/pkg/tool/gomemoryd/store"

func (s *Service) ListRelated(identifier int64) ([]store.Related, error) {
	return s.store.ListRelated(identifier)
}
