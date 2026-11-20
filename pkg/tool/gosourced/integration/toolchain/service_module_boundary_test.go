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
	o := service_tester.New(
		t,
		service_tester.WithFixture("move-cgo-reference/src"),
	)
	r, e := o.Symbol.MoveSymbol(
		t.Context(),
		o.Directory,
		"example/pkg/target",
		"ItemFields",
		"example/pkg/target/constant",
		"constant.go",
		true,
		true,
	)
	assert.FatalOnError(t, e)
	rewrite_tester.New(t).ModuleEntries(r)
}

func TestMovePackageWithTestVariantsStaysInsideModule(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("test-home/src"))
	e := os.MkdirAll(filepath.Join(o.Directory, "pkg/alfa/integration"), 0755)
	assert.FatalOnError(t, e)
	r, f := o.Packages.MovePackage(
		t.Context(),
		o.Directory,
		"example/pkg/alfa/integration_test/client",
		"example/pkg/alfa/integration/client",
		false,
	)
	assert.FatalOnError(t, f)
	rewrite_tester.New(t).ModuleEntries(r)
}
