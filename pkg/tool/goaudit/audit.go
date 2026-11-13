package goaudit

import (
	"github.com/funtimecoding/soil/pkg/lint"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/warm_tree"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/constant"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/scan"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/scan/audit_configuration"
	"path/filepath"
)

func Audit(
	t *warm_tree.Tree,
	configuration *audit_configuration.Configuration,
) *output.Results {
	v := scan.Tree(t)
	services := scan.Services(v, filepath.Base(t.Root()), configuration)
	r := headless(v, services, scan.IdentityWarnings(v, services))
	r.Header = lint.HeaderLine(constant.Identity.Name(), t.Root(), "")

	return r
}
