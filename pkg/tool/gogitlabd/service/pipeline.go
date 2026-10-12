package service

import "github.com/funtimecoding/soil/pkg/gitlab/pipeline_detail"

func (s *Service) Pipeline(
	project int64,
	identifier int64,
) (*pipeline_detail.Detail, error) {
	return s.client.Pipeline(project, identifier)
}
