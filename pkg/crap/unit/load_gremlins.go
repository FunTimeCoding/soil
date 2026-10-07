package unit

import (
	"github.com/funtimecoding/soil/pkg/crap/constant"
	"github.com/funtimecoding/soil/pkg/crap/mutation_report"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"path/filepath"
	"testing"
)

func loadGremlins(t *testing.T) *mutation_report.Report {
	t.Helper()
	root := t.TempDir()
	testutil.WriteFile(t, root, "report.json", constant.FixtureGremlins)

	return mutation_report.Load(filepath.Join(root, "report.json"))
}
