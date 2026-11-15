package module_tester

import "github.com/funtimecoding/soil/pkg/system/system_tester"

func (o *Module) File(
	path string,
	content string,
) *Module {
	o.t.Helper()
	system_tester.Write(o.t, o.directory, path, content)

	return o
}
