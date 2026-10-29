package reporter

import (
	"github.com/funtimecoding/soil/pkg/errors/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func New(project string) *Reporter {
	return build(project, environment.Required(constant.LocatorEnvironment))
}
