package security_tester

import "github.com/funtimecoding/soil/pkg/tool/gosecurityd/service/overview"

func (o *Tester) Overview() *overview.Service {
	o.t.Helper()

	return overview.New(o.Flag(), o.Debian())
}
