package face

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/variable"
)

type VariableForge interface {
	Variables(
		x context.Context,
		project int64,
	) ([]*variable.Variable, error)
	ProjectVariable(
		x context.Context,
		project int64,
		key string,
	) (*variable.Variable, error)
	CreateProjectVariable(
		x context.Context,
		project int64,
		key string,
		value string,
		protected bool,
		masked bool,
		literal bool,
	) (*variable.Variable, error)
	UpdateProjectVariable(
		x context.Context,
		project int64,
		key string,
		value string,
		protected bool,
		masked bool,
		literal bool,
	) (*variable.Variable, error)
	DeleteProjectVariable(
		x context.Context,
		project int64,
		key string,
	) error
}
