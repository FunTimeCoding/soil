package unit

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"os"
	"path/filepath"
	"testing"
)

func prepareReplacedLibrary(
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
	e := os.WriteFile(
		filepath.Join(library, "go.mod"),
		[]byte("module other.test/lib\n\ngo 1.26.3\n"),
		0644,
	)

	if e != nil {
		t.Fatalf("write go.mod: %s", e)
	}

	user := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata(filepath.Join(fixture, "user", "src")),
		"require other.test/lib v0.0.0",
		fmt.Sprintf("replace other.test/lib => %s", library),
	)

	return library, user
}
