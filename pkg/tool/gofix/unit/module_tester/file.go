package module_tester

import "github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"

func (o *Module) File(
	path string,
	content string,
) *Module {
	o.t.Helper()
	testutil.WriteFile(o.t, o.directory, path, content)

	return o
}
