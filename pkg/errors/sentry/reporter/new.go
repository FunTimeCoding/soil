package reporter

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func New(
	project string,
	version string,
) *Reporter {
	return &Reporter{
		project: project,
		locator: environment.Required(constant.LocatorEnvironment),
		version: version,
	}
}
