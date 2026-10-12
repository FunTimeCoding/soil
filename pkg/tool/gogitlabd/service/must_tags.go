package service

import "github.com/funtimecoding/soil/pkg/gitlab/tag"

func (s *Service) MustTags(project int64) []*tag.Tag {
	return s.client.MustTags(project)
}
