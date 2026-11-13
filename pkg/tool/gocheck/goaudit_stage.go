package gocheck

import (
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/goaudit"
	goauditConstant "github.com/funtimecoding/soil/pkg/tool/goaudit/constant"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/scan/audit_configuration"
	"github.com/funtimecoding/soil/pkg/tool/gocheck/option"
	"path/filepath"
)

func goauditStage(o *option.Check) []*output.Results {
	var paths []string

	for _, p := range goauditConstant.ConfigurationPaths {
		paths = append(paths, filepath.Join(o.Root, p))
	}

	configuration := audit_configuration.Load(system.FirstFile(paths...))

	return []*output.Results{
		goaudit.Audit(o.Trees.Tree(o.Root), configuration),
		goaudit.Permissions(o.Root, configuration),
	}
}
