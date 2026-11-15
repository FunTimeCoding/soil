package rewrite_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"testing"
)

func IndexedAndFull(
	t *testing.T,
) (string, *service_tester.Services, *service_tester.Services) {
	t.Helper()
	directory := service_tester.Prepare(t, "indexed/src")
	full := service_tester.NewService(t)
	full.Workspace.UseFullLoad()

	return directory, service_tester.NewService(t), full
}
