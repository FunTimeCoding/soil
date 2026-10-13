package service

import "github.com/funtimecoding/soil/pkg/gitlab/variable"

func (s *Service) UpdateProjectVariable(
	instance string,
	project int64,
	key string,
	value string,
	protected bool,
	masked bool,
	literal bool,
) (*variable.Variable, error) {
	c, e := s.pool.Client(instance)

	if e != nil {
		return nil, e
	}

	return c.UpdateProjectVariable(
		project,
		key,
		value,
		protected,
		masked,
		literal,
	)
}
