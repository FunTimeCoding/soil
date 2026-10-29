package reporter

import (
	"github.com/funtimecoding/soil/pkg/errors/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func NewOptional(project string) *Reporter {
	return build(project, environment.Optional(constant.LocatorEnvironment))
}
