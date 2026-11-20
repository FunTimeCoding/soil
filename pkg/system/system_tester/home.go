package system_tester

import soil "github.com/funtimecoding/soil/pkg/constant"

func (o *Tester) Home() string {
	o.t.Helper()
	result := o.t.TempDir()
	o.t.Setenv(soil.HomeEnvironment, result)

	return result
}
