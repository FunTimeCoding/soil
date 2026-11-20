package module

import "github.com/funtimecoding/soil/pkg/system/system_tester"

func (o *Module) File(
	path string,
	content string,
) *Module {
	o.t.Helper()
	system_tester.New(o.t).Write(o.directory, path, content)

	return o
}
