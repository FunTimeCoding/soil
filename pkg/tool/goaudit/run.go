package goaudit

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/lint"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/warm_tree"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/constant"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/option"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/scan"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/scan/audit_configuration"
	"os"
	"path/filepath"
)

func Run(o *option.Audit) {
	configuration := audit_configuration.Load(
		system.FirstFile(constant.ConfigurationPaths...),
	)
	configuration.Daemon = configuration.Daemon || o.Daemon
	failed := false

	for _, root := range o.Roots {
		if !o.Web && !o.Table {
			if output.Print(Audit(warm_tree.New(root), configuration), false) {
				failed = true
			}

			continue
		}

		lint.Header(constant.Identity.Name(), root, "")
		v := scan.Tree(warm_tree.New(root))
		repo := filepath.Base(root)
		services := scan.Services(v, repo, configuration)

		if o.Web {
			runWeb(scan.Frontends(v, services))

			continue
		}

		runTable(
			services,
			scan.IdentityWarnings(v, services),
			scan.Clients(v, repo, configuration),
		)
	}

	if !o.Web && !o.Table &&
		output.PrintResults(
			Permissions(soil.CurrentDirectory, configuration).Entries,
			false,
		) {
		failed = true
	}

	if failed {
		os.Exit(1)
	}
}
