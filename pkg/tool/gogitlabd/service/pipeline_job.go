package service

import "github.com/funtimecoding/soil/pkg/gitlab/job"

func (s *Service) PipelineJob(
	project int64,
	identifier int64,
) (*job.Job, error) {
	return s.client.PipelineJob(project, identifier)
}
