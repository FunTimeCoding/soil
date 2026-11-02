package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/face"
	"github.com/funtimecoding/soil/pkg/gitlab/job"
)

func (s *Service) PipelineJob(
	x context.Context,
	instance string,
	project int64,
	identifier int64,
) (*job.Job, error) {
	return guarded(
		x,
		s,
		instance,
		func(y context.Context, c face.GitLabClient) (*job.Job, error) {
			return c.PipelineJob(y, project, identifier)
		},
	)
}
