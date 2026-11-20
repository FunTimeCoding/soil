package module_tester

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/source/inventory"
	"github.com/funtimecoding/soil/pkg/system/system_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"path/filepath"
)

func (o *Tester) ReplacedLibrary(
	fixture string,
) (string, *service_tester.Tester) {
	o.t.Helper()
	library := o.t.TempDir()
	testutil.CopyTree(
		o.t,
		service_tester.ServiceTestdata(filepath.Join(fixture, "lib")),
		library,
	)
	system_tester.New(o.t).Write(
		library,
		"go.mod",
		"module other.test/lib\n\ngo 1.26.3\n",
	)
	user := testutil.PrepareTestPackage(
		o.t,
		service_tester.ServiceTestdata(filepath.Join(fixture, "user", "src")),
		"require other.test/lib v0.0.0",
		fmt.Sprintf("replace other.test/lib => %s", library),
	)
	i := inventory.New()
	i.Add("lib", library)
	i.Add("user", user)
	result := service_tester.New(
		o.t,
		service_tester.WithDirectory(user),
		service_tester.WithInventory(i),
	)
	result.Workspace.UseIndex(service_tester.IndexDirectory())

	return library, result
}
