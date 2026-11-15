package usage_tester

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/system/system_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"testing"
)

func Prepare(t *testing.T) string {
	t.Helper()
	gorm := t.TempDir()
	testutil.CopyTree(
		t,
		service_tester.ServiceTestdata("service-usage/gorm"),
		gorm,
	)
	system_tester.Write(t, gorm, "go.mod", "module gorm.io/gorm\n\ngo 1.26.3\n")

	return service_tester.Prepare(
		t,
		"service-usage/src",
		"require gorm.io/gorm v0.0.0",
		fmt.Sprintf("replace gorm.io/gorm => %s", gorm),
	)
}
