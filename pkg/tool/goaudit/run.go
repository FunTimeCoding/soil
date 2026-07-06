package goaudit

import (
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/option"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/scan"
	"path/filepath"
)

func Run(o *option.Audit) {
	for _, root := range o.Roots {
		v := virtual_file_system.From(root)
		repo := filepath.Base(root)
		configuration := scan.LoadConfiguration(".goaudit.yaml")
		services := scan.Services(v, repo, configuration)
		identityWarnings := scan.IdentityWarnings(v, services)
		clients := scan.Clients(v, repo, configuration)

		if o.Table {
			runTable(services, identityWarnings, clients)

			continue
		}

		runHeadless(v, services, identityWarnings)
	}
}
