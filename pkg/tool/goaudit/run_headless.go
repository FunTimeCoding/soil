package goaudit

import (
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/scan"
)

func runHeadless(
	v *virtual_file_system.System,
	e []*scan.Service,
	identityWarnings []*concern.Concern,
) bool {
	r := output.NewResults()

	for _, s := range e {
		for _, c := range s.Concerns {
			r.AddConcern(c)
		}
	}

	for _, c := range identityWarnings {
		r.AddConcern(c)
	}

	for _, c := range scan.MissingSentry(v) {
		r.AddConcern(c)
	}

	for _, c := range scan.MisplacedTests(v) {
		r.AddConcern(c)
	}

	for _, c := range scan.ConstantPlacement(v) {
		r.AddConcern(c)
	}

	return output.PrintResults(r.Entries, false)
}
