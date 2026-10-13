package service

import "github.com/funtimecoding/soil/pkg/gitlab/job"

func (s *Service) CancelJob(
	instance string,
	project int64,
	identifier int64,
) (*job.Job, error) {
	c, e := s.pool.Client(instance)

	if e != nil {
		return nil, e
	}

	return c.CancelJob(project, identifier)
}
