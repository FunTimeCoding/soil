package module_tester

import (
	"github.com/funtimecoding/soil/pkg/system/system_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"testing"
)

func DottedModule(
	t *testing.T,
	fixture string,
) *service_tester.Tester {
	t.Helper()
	result := service_tester.New(t, service_tester.WithFixture(fixture))
	system_tester.Write(
		t,
		result.Directory,
		"go.mod",
		"module example.org/m\n\ngo 1.26.3\n",
	)

	return result
}
