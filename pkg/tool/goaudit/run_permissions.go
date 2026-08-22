package goaudit

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/scan"
)

func runPermissions(configuration *scan.Configuration) bool {
	r := output.NewResults()

	for _, c := range scan.ModelContextPermissions(
		constant.CurrentDirectory,
		configuration,
	) {
		r.AddConcern(c)
	}

	return output.PrintResults(r.Entries, false)
}
