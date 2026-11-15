package module_tester

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/system/system_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"path/filepath"
	"testing"
)

func ReplacedLibrary(
	t *testing.T,
	fixture string,
) (string, string) {
	t.Helper()
	library := t.TempDir()
	testutil.CopyTree(
		t,
		service_tester.ServiceTestdata(filepath.Join(fixture, "lib")),
		library,
	)
	system_tester.Write(
		t,
		library,
		"go.mod",
		"module other.test/lib\n\ngo 1.26.3\n",
	)
	user := service_tester.Prepare(
		t,
		filepath.Join(fixture, "user", "src"),
		"require other.test/lib v0.0.0",
		fmt.Sprintf("replace other.test/lib => %s", library),
	)

	return library, user
}
