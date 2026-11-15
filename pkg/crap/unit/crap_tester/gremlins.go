package crap_tester

import (
	"github.com/funtimecoding/soil/pkg/crap/constant"
	"github.com/funtimecoding/soil/pkg/crap/mutation_report"
	"github.com/funtimecoding/soil/pkg/system/system_tester"
	"path/filepath"
	"testing"
)

func Gremlins(t *testing.T) *mutation_report.Report {
	t.Helper()
	root := t.TempDir()
	system_tester.Write(t, root, "report.json", constant.FixtureGremlins)

	return mutation_report.Load(filepath.Join(root, "report.json"))
}
