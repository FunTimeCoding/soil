package service

import (
	"github.com/funtimecoding/soil/pkg/gitlab/face"
	"github.com/funtimecoding/soil/pkg/gitlab/job"
)

func (s *Service) PipelineJob(
	instance string,
	project int64,
	identifier int64,
) (*job.Job, error) {
	return guarded(
		s,
		instance,
		func(c face.GitLabClient) (*job.Job, error) {
			return c.PipelineJob(project, identifier)
		},
	)
}
