package goaudit

import (
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/constant"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/option"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/scan"
	"os"
	"path/filepath"
)

func Run(o *option.Audit) {
	configuration := scan.LoadConfiguration(
		system.FirstFile(constant.ConfigurationPaths...),
	)
	failed := false

	for _, root := range o.Roots {
		v := virtual_file_system.From(root)
		repo := filepath.Base(root)
		services := scan.Services(v, repo, configuration)

		if o.Web {
			runWeb(scan.Frontends(v, services))

			continue
		}

		identityWarnings := scan.IdentityWarnings(v, services)
		clients := scan.Clients(v, repo, configuration)

		if o.Table {
			runTable(services, identityWarnings, clients)

			continue
		}

		if runHeadless(v, services, identityWarnings) {
			failed = true
		}
	}

	if !o.Web && !o.Table && runPermissions(configuration) {
		failed = true
	}

	if failed {
		os.Exit(1)
	}
}
