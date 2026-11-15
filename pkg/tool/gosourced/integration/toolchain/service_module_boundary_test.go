package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/integration/toolchain/rewrite_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"os"
	"path/filepath"
	"testing"
)

func TestMoveWithCgoConsumerStaysInsideModule(t *testing.T) {
	d := service_tester.Prepare(t, "move-cgo-reference/src")
	s := service_tester.NewService(t)
	r, e := s.Symbol.MoveSymbol(
		t.Context(),
		d,
		"example/pkg/target",
		"ItemFields",
		"example/pkg/target/constant",
		"constant.go",
		true,
		true,
	)
	assert.FatalOnError(t, e)
	rewrite_tester.ModuleEntries(t, r)
}

func TestMovePackageWithTestVariantsStaysInsideModule(t *testing.T) {
	d := service_tester.Prepare(t, "test-home/src")
	e := os.MkdirAll(filepath.Join(d, "pkg/alfa/integration"), 0755)
	assert.FatalOnError(t, e)
	s := service_tester.NewService(t)
	r, f := s.Packages.MovePackage(
		t.Context(),
		d,
		"example/pkg/alfa/integration_test/client",
		"example/pkg/alfa/integration/client",
		false,
	)
	assert.FatalOnError(t, f)
	rewrite_tester.ModuleEntries(t, r)
}
