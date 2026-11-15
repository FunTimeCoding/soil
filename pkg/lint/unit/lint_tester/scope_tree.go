package lint_tester

import (
	"github.com/funtimecoding/soil/pkg/system/system_tester"
	"testing"
)

func ScopeTree(t *testing.T) string {
	t.Helper()
	root := t.TempDir()
	system_tester.Write(t, root, "doc/ai/spec/naming.md", "# Naming\n")
	system_tester.Write(t, root, "pkg/lint/lint.go", "package lint\n")

	return root
}
