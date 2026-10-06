package unit

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"testing"
)

func indexedAndFull(t *testing.T) (string, *service.Service, *service.Service) {
	t.Helper()
	directory := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("indexed/src"),
	)
	full := testService()
	full.UseFullLoad()

	return directory, testService(), full
}
