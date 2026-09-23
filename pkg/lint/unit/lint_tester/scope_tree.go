package lint_tester

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"testing"
)

func ScopeTree(t *testing.T) string {
	t.Helper()
	root := t.TempDir()
	testutil.WriteFile(t, root, "doc/ai/spec/naming.md", "# Naming\n")
	testutil.WriteFile(t, root, "pkg/lint/lint.go", "package lint\n")

	return root
}
