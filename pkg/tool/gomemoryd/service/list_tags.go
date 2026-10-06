package service

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"

func (s *Service) ListTags() ([]record.TagCount, error) {
	return s.store.ListTags()
}
