package crap_tester

import (
	"github.com/funtimecoding/soil/pkg/crap/constant"
	"github.com/funtimecoding/soil/pkg/crap/mutation_report"
	"github.com/funtimecoding/soil/pkg/system/system_tester"
	"path/filepath"
)

func (o *Tester) Gremlins() *mutation_report.Report {
	o.t.Helper()
	root := o.t.TempDir()
	system_tester.New(o.t).Write(root, "report.json", constant.FixtureGremlins)

	return mutation_report.Load(filepath.Join(root, "report.json"))
}
