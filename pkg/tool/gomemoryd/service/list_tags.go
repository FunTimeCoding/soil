package service

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/store"

func (s *Service) ListTags() ([]store.TagCount, error) {
	return s.store.ListTags()
}
