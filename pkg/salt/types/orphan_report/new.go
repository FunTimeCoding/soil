package orphan_report

import "github.com/funtimecoding/soil/pkg/salt/types/analysis"

func New(a *analysis.Analysis) *Report {
	return &Report{
		Summary: &Summary{
			Reached:        len(a.Reached),
			Possible:       len(a.Possible),
			Unreached:      len(a.Unreached),
			Missing:        len(a.Missing),
			UnreachedFiles: len(a.UnreachedFiles),
		},
		Unreached:        a.Unreached,
		UnreachedFiles:   a.UnreachedFiles,
		Possible:         a.Possible,
		Missing:          a.Missing,
		DeadEnvironments: a.DeadEnvironments,
	}
}
