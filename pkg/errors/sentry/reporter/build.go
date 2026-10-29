package reporter

import (
	"github.com/funtimecoding/soil/pkg/errors/constant"
	"github.com/funtimecoding/soil/pkg/stamp"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/web/stall_client"
	"github.com/mattn/go-isatty"
	"os"
)

func build(
	project string,
	locator string,
) *Reporter {
	s := stamp.New()

	return &Reporter{
		project: project,
		locator: locator,
		version: s.Tag(),
		module:  s.Module,
		echo:    environment.Exists(constant.PrintEnvironment),
		output:  os.Stderr,
		color:   isatty.IsTerminal(os.Stderr.Fd()),
		pool:    stall_client.New().Build(),
	}
}
