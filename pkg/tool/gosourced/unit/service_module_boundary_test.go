package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"os"
	"path/filepath"
	"testing"
)

func TestMoveWithCgoConsumerStaysInsideModule(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("move-cgo-reference/src"),
	)
	s := testService()
	r, e := s.MoveSymbol(
		d,
		"example/pkg/target",
		"ItemFields",
		"example/pkg/target/constant",
		"constant.go",
		true,
		true,
	)
	assert.FatalOnError(t, e)
	assertEntriesInsideModule(t, r)
}

func TestMovePackageWithTestVariantsStaysInsideModule(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("test-home/src"),
	)
	e := os.MkdirAll(filepath.Join(d, "pkg/alfa/integration"), 0755)
	assert.FatalOnError(t, e)
	s := testService()
	r, f := s.MovePackage(
		d,
		"example/pkg/alfa/integration_test/client",
		"example/pkg/alfa/integration/client",
		false,
	)
	assert.FatalOnError(t, f)
	assertEntriesInsideModule(t, r)
}
