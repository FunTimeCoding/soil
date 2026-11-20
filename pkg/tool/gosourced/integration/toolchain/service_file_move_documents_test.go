package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/system/system_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"testing"
)

func TestMoveSymbolsRetargetsEmptiedFile(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("move-visibility/src"),
	)
	system_tester.New(t).Write(
		o.Directory,
		"doc/note.md",
		"Keep `pkg/alfa/keep.go`, move `pkg/alfa/lone.go`.\n",
	)
	r, e := o.Symbol.MoveSymbols(
		t.Context(),
		o.Directory,
		"example/pkg/alfa",
		[]string{"lone"},
		"",
		"example/pkg/bravo",
		"",
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.String(
		t,
		"Keep `pkg/alfa/keep.go`, move `pkg/bravo/lone.go`.\n",
		o.Read("doc/note.md"),
	)
}

func TestMoveSymbolsRetargetsSplitFileToPackage(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("move-split-file/src"),
	)
	system_tester.New(t).Write(
		o.Directory,
		"doc/note.md",
		"Helpers live in `pkg/alfa/helpers.go`.\n",
	)
	r, e := o.Symbol.MoveSymbols(
		t.Context(),
		o.Directory,
		"example/pkg/alfa",
		[]string{"Clamp", "Limit"},
		"",
		"example/pkg/bravo",
		"",
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.String(t, "Helpers live in `pkg/bravo`.\n", o.Read("doc/note.md"))
}

func TestExtractToFileRetargetsRenamedResidual(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("extract-last-pair/src"),
	)
	system_tester.New(t).Write(
		o.Directory,
		"doc/note.md",
		"Names live in `pkg/target/combined.go`.\n",
	)
	r, e := o.File.ExtractToFile(
		t.Context(),
		o.Directory,
		"pkg/target/combined.go",
		"FormatName",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.String(
		t,
		"Names live in `pkg/target/summarize_name.go`.\n",
		o.Read("doc/note.md"),
	)
}
