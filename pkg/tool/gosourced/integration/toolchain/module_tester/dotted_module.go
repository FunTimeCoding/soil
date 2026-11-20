package module_tester

import (
	"github.com/funtimecoding/soil/pkg/system/system_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
)

func (o *Tester) DottedModule(fixture string) *service_tester.Tester {
	o.t.Helper()
	result := service_tester.New(o.t, service_tester.WithFixture(fixture))
	system_tester.New(o.t).Write(
		result.Directory,
		"go.mod",
		"module example.org/m\n\ngo 1.26.3\n",
	)

	return result
}
