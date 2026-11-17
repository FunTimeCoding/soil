package rewrite_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"testing"
)

func IndexedAndFull(
	t *testing.T,
) (*service_tester.Tester, *service_tester.Services) {
	t.Helper()
	o := service_tester.New(t, service_tester.WithFixture("indexed/src"))
	full := service_tester.New(t).Services
	full.Workspace.UseFullLoad()

	return o, full
}
