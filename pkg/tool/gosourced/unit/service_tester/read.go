package service_tester

import "github.com/funtimecoding/soil/pkg/system/system_tester"

func (o *Tester) Read(path string) string {
	o.t.Helper()

	return system_tester.New(o.t).Read(o.Directory, path)
}
