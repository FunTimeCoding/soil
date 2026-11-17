package service_tester

import "github.com/funtimecoding/soil/pkg/system/system_tester"

func (o *Tester) Read(path string) string {
	o.t.Helper()

	return system_tester.Read(o.t, o.Directory, path)
}
