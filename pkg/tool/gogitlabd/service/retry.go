package service

import "github.com/funtimecoding/soil/pkg/gitlab/job"

func (s *Service) Retry(
	project int64,
	jobIdentifier int64,
) (*job.Job, error) {
	return s.client.Retry(project, jobIdentifier)
}
