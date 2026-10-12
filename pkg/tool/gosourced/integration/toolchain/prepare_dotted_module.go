package toolchain

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"os"
	"path/filepath"
	"testing"
)

func prepareDottedModule(
	t *testing.T,
	fixture string,
) string {
	t.Helper()
	d := testutil.PrepareTestPackage(t, service_tester.ServiceTestdata(fixture))
	e := os.WriteFile(
		filepath.Join(d, "go.mod"),
		[]byte("module example.org/m\n\ngo 1.26.3\n"),
		0644,
	)

	if e != nil {
		t.Fatalf("write go.mod: %s", e)
	}

	return d
}
