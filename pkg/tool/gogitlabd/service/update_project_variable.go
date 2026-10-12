package service

import "github.com/funtimecoding/soil/pkg/gitlab/variable"

func (s *Service) UpdateProjectVariable(
	project int64,
	key string,
	value string,
	protected bool,
	masked bool,
	literal bool,
) (*variable.Variable, error) {
	return s.client.UpdateProjectVariable(
		project,
		key,
		value,
		protected,
		masked,
		literal,
	)
}
