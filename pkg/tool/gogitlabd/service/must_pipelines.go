package service

import "github.com/funtimecoding/soil/pkg/gitlab/pipeline"

func (s *Service) MustPipelines(project int64) []*pipeline.Pipeline {
	return s.client.MustPipelines(project)
}
