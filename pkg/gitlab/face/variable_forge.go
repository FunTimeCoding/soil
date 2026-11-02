package face

import "github.com/funtimecoding/soil/pkg/gitlab/variable"

type VariableForge interface {
	Variables(project int64) ([]*variable.Variable, error)
	ProjectVariable(
		project int64,
		key string,
	) (*variable.Variable, error)
	CreateProjectVariable(
		project int64,
		key string,
		value string,
		protected bool,
		masked bool,
		literal bool,
	) (*variable.Variable, error)
	UpdateProjectVariable(
		project int64,
		key string,
		value string,
		protected bool,
		masked bool,
		literal bool,
	) (*variable.Variable, error)
	DeleteProjectVariable(
		project int64,
		key string,
	) error
}
