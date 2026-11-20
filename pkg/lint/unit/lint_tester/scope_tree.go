package lint_tester

import "github.com/funtimecoding/soil/pkg/system/system_tester"

func (o *Tester) ScopeTree() string {
	o.t.Helper()
	root := o.t.TempDir()
	system_tester.New(o.t).Write(root, "doc/ai/spec/naming.md", "# Naming\n")
	system_tester.New(o.t).Write(root, "pkg/lint/lint.go", "package lint\n")

	return root
}
