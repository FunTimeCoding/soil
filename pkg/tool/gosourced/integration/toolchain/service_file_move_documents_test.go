package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"testing"
)

func TestMoveSymbolsRetargetsAnEmptiedFile(t *testing.T) {
	d := moveVisibilityFixture(t)
	testutil.WriteFile(
		t,
		d,
		"doc/note.md",
		"Keep `pkg/alfa/keep.go`, move `pkg/alfa/lone.go`.\n",
	)
	r, e := service_tester.NewService(t).MoveSymbols(
		t.Context(),
		d,
		"example/pkg/alfa",
		[]string{"lone"},
		"",
		"example/pkg/bravo",
		"",
		false,
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.String(
		t,
		"Keep `pkg/alfa/keep.go`, move `pkg/bravo/lone.go`.\n",
		service_tester.ReadFixtureFile(t, d, "doc/note.md"),
	)
}

func TestExtractToFileRetargetsTheRenamedResidual(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("extract-last-pair/src"),
	)
	testutil.WriteFile(
		t,
		d,
		"doc/note.md",
		"Names live in `pkg/target/combined.go`.\n",
	)
	r, e := service_tester.NewService(t).ExtractToFile(
		t.Context(),
		d,
		"pkg/target/combined.go",
		"FormatName",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.String(
		t,
		"Names live in `pkg/target/summarize_name.go`.\n",
		service_tester.ReadFixtureFile(t, d, "doc/note.md"),
	)
}
