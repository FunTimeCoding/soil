package service

import (
	"github.com/funtimecoding/soil/pkg/gitlab/face"
	"github.com/funtimecoding/soil/pkg/gitlab/variable"
)

func (s *Service) CreateProjectVariable(
	instance string,
	project int64,
	key string,
	value string,
	protected bool,
	masked bool,
	literal bool,
) (*variable.Variable, error) {
	return guarded(
		s,
		instance,
		func(c face.GitLabClient) (*variable.Variable, error) {
			return c.CreateProjectVariable(
				project,
				key,
				value,
				protected,
				masked,
				literal,
			)
		},
	)
}
