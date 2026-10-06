package reporter

import (
	"github.com/funtimecoding/soil/pkg/errors/constant"
	"github.com/funtimecoding/soil/pkg/stamp"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func New(project string) *Reporter {
	return &Reporter{
		project: project,
		locator: environment.Required(constant.LocatorEnvironment),
		version: stamp.New().Tag(),
	}
}
