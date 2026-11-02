package service

import (
	"github.com/funtimecoding/soil/pkg/gitlab/face"
	"github.com/funtimecoding/soil/pkg/gitlab/job"
)

func (s *Service) Retry(
	instance string,
	project int64,
	jobIdentifier int64,
) (*job.Job, error) {
	return guarded(
		s,
		instance,
		func(c face.GitLabClient) (*job.Job, error) {
			return c.Retry(project, jobIdentifier)
		},
	)
}
