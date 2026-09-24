package unit

import (
	"github.com/funtimecoding/soil/pkg/crap/constant"
	"github.com/funtimecoding/soil/pkg/crap/mutation"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"path/filepath"
	"testing"
)

func loadGremlins(t *testing.T) *mutation.Report {
	t.Helper()
	root := t.TempDir()
	testutil.WriteFile(t, root, "report.json", constant.FixtureGremlins)

	return mutation.Load(filepath.Join(root, "report.json"))
}
