package break_tester

import (
	gitConstant "github.com/funtimecoding/soil/pkg/git/constant"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/system_tester"
	"github.com/funtimecoding/soil/pkg/tool/gobreak/constant"
	"path/filepath"
	"testing"
)

func Fixture(
	t *testing.T,
	breaks string,
) string {
	t.Helper()
	t.Setenv("GOFLAGS", "-buildvcs=false")
	root := t.TempDir()
	system.MakeDirectory(filepath.Join(root, gitConstant.Directory))
	system_tester.New(t).Write(root, "go.mod", constant.FixtureModule)
	system_tester.New(t).Write(root, "shape/shape.go", constant.FixtureShape)
	system_tester.New(t).Write(
		root,
		"shape/shape_test.go",
		constant.FixtureShapeTest,
	)
	system_tester.New(t).Write(root, "probe/breaks.yaml", breaks)

	return filepath.Join(root, "probe", "breaks.yaml")
}
