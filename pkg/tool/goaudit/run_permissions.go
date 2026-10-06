package goaudit

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/scan"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/scan/audit_configuration"
)

func runPermissions(configuration *audit_configuration.Configuration) bool {
	r := output.NewResults()

	for _, c := range scan.ModelContextPermissions(
		constant.CurrentDirectory,
		configuration,
	) {
		r.AddConcern(c)
	}

	return output.PrintResults(r.Entries, false)
}
