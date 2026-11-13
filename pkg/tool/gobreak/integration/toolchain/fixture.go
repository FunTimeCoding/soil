package toolchain

import (
	gitConstant "github.com/funtimecoding/soil/pkg/git/constant"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/gobreak/constant"
	"path/filepath"
	"testing"
)

func fixture(
	t *testing.T,
	breaks string,
) string {
	t.Helper()
	t.Setenv("GOFLAGS", "-buildvcs=false")
	root := t.TempDir()
	system.MakeDirectory(filepath.Join(root, gitConstant.Directory))
	write(root, "go.mod", constant.FixtureModule)
	write(root, "shape/shape.go", constant.FixtureShape)
	write(root, "shape/shape_test.go", constant.FixtureShapeTest)
	write(root, "probe/breaks.yaml", breaks)

	return filepath.Join(root, "probe", "breaks.yaml")
}
