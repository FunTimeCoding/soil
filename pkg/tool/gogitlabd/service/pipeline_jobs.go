package service

import "github.com/funtimecoding/soil/pkg/gitlab/job"

func (s *Service) PipelineJobs(
	instance string,
	project int64,
	identifier int64,
) ([]*job.Job, error) {
	c, e := s.pool.Client(instance)

	if e != nil {
		return nil, e
	}

	return c.PipelineJobs(project, identifier)
}
