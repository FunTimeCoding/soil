package goaudit

import (
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/scan"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/scan/audit_configuration"
)

func Permissions(
	base string,
	configuration *audit_configuration.Configuration,
) *output.Results {
	r := output.NewResults()

	for _, c := range scan.ModelContextPermissions(base, configuration) {
		r.AddConcern(c)
	}

	return r
}
