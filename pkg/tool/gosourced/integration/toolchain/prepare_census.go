package toolchain

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"os"
	"path/filepath"
	"testing"
)

func prepareCensus(t *testing.T) string {
	t.Helper()
	gorm := t.TempDir()
	testutil.CopyTree(
		t,
		service_tester.ServiceTestdata(filepath.Join("service-census", "gorm")),
		gorm,
	)
	e := os.WriteFile(
		filepath.Join(gorm, "go.mod"),
		[]byte("module gorm.io/gorm\n\ngo 1.26.3\n"),
		0644,
	)

	if e != nil {
		t.Fatalf("write go.mod: %s", e)
	}

	return testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata(filepath.Join("service-census", "src")),
		"require gorm.io/gorm v0.0.0",
		fmt.Sprintf("replace gorm.io/gorm => %s", gorm),
	)
}
