package rewrite_tester

import "github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"

func (o *Tester) IndexedAndFull() (
	*service_tester.Tester,
	*service_tester.Services,
) {
	o.t.Helper()
	s := service_tester.New(o.t, service_tester.WithFixture("indexed/src"))
	full := service_tester.New(o.t).Services
	full.Workspace.UseFullLoad()

	return s, full
}
