package reporter

import (
	"github.com/funtimecoding/soil/pkg/errors/constant"
	"github.com/funtimecoding/soil/pkg/stamp"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/web/stall_client"
)

func New(project string) *Reporter {
	return &Reporter{
		project: project,
		locator: environment.Required(constant.LocatorEnvironment),
		version: stamp.New().Tag(),
		pool:    stall_client.New().Build(),
	}
}
