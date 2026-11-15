package unit

import (
	"bytes"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/system"
	textConstant "github.com/funtimecoding/soil/pkg/text/constant"
	"github.com/funtimecoding/soil/pkg/tool/goreplace"
	"github.com/funtimecoding/soil/pkg/tool/goreplace/unit/goreplace_tester"
	"os"
	"testing"
)

func TestReplaceAppliesEveryBlock(t *testing.T) {
	path := goreplace_tester.Target(t, "alpha\nbeta\ngamma\n")
	var out bytes.Buffer
	assert.FatalOnError(
		t,
		goreplace.Run(
			path,
			goreplace_tester.Blocks("alpha", "ALPHA", "gamma", "GAMMA"),
			false,
			&out,
		),
	)
	assert.String(t, "ALPHA\nbeta\nGAMMA\n", system.ReadFileUnsafe(path))
	assert.StringContains(t, "@@ block 2, line 3\n-gamma\n+GAMMA", out.String())
	assert.StringContains(t, "2 blocks applied", out.String())
}

func TestReplaceDryRunWritesNothing(t *testing.T) {
	path := goreplace_tester.Target(t, "alpha\nbeta\n")
	var out bytes.Buffer
	assert.FatalOnError(
		t,
		goreplace.Run(
			path,
			goreplace_tester.Blocks("alpha", "ALPHA"),
			true,
			&out,
		),
	)
	assert.String(t, "alpha\nbeta\n", system.ReadFileUnsafe(path))
	assert.StringContains(t, "would apply (dry run", out.String())
}

func TestReplaceWritesNothingWhenOneBlockFails(t *testing.T) {
	path := goreplace_tester.Target(t, "alpha\nbeta\n")
	e := goreplace.Run(
		path,
		goreplace_tester.Blocks("alpha", "ALPHA", "missing", "x"),
		false,
		&bytes.Buffer{},
	)
	assert.StringContains(t, "block 2: search not found", assert.Failure(t, e))
	assert.String(t, "alpha\nbeta\n", system.ReadFileUnsafe(path))
}

func TestReplaceRefusesAmbiguousSearch(t *testing.T) {
	path := goreplace_tester.Target(t, "x\ny\nx\n")
	e := goreplace.Run(
		path,
		goreplace_tester.Blocks("x", "z"),
		false,
		&bytes.Buffer{},
	)
	assert.StringContains(
		t,
		"matches 2 times, at lines 1, 3",
		assert.Failure(t, e),
	)
}

func TestReplaceRefusesOverlappingBlocks(t *testing.T) {
	path := goreplace_tester.Target(t, "alpha\nbeta\ngamma\n")
	e := goreplace.Run(
		path,
		goreplace_tester.Blocks("alpha\nbeta", "x", "beta\ngamma", "y"),
		false,
		&bytes.Buffer{},
	)
	assert.StringContains(
		t,
		"blocks 1 and 2 overlap at line 2",
		assert.Failure(t, e),
	)
}

func TestReplaceMatchesAgainstOriginal(t *testing.T) {
	path := goreplace_tester.Target(t, "alpha\nbeta\n")
	e := goreplace.Run(
		path,
		goreplace_tester.Blocks("alpha", "delta", "delta", "epsilon"),
		false,
		&bytes.Buffer{},
	)
	assert.StringContains(t, "block 2: search not found", assert.Failure(t, e))
}

func TestReplaceDeletesWholeLine(t *testing.T) {
	path := goreplace_tester.Target(t, "alpha\nbeta\ngamma\n")
	assert.FatalOnError(
		t,
		goreplace.Run(
			path,
			goreplace_tester.Blocks("beta\n", ""),
			false,
			&bytes.Buffer{},
		),
	)
	assert.String(t, "alpha\ngamma\n", system.ReadFileUnsafe(path))
}

func TestReplaceHintsAtWhitespace(t *testing.T) {
	path := goreplace_tester.Target(t, "  indented line\n")
	e := goreplace.Run(
		path,
		goreplace_tester.Blocks("indented  line", "x"),
		false,
		&bytes.Buffer{},
	)
	assert.StringContains(
		t,
		"matches when whitespace is ignored",
		assert.Failure(t, e),
	)
}

func TestReplaceHintsAtFirstLine(t *testing.T) {
	path := goreplace_tester.Target(t, "one\ntwo\n")
	e := goreplace.Run(
		path,
		goreplace_tester.Blocks("one\nthree", "x"),
		false,
		&bytes.Buffer{},
	)
	assert.StringContains(
		t,
		"its first line occurs at line 1",
		assert.Failure(t, e),
	)
}

func TestReplaceRefusesStrayText(t *testing.T) {
	path := goreplace_tester.Target(t, "alpha\n")
	e := goreplace.Run(path, "hello\n", false, &bytes.Buffer{})
	assert.StringContains(
		t,
		`line 1: expected <<<<<<< SEARCH, SPAN, AFTER or BEFORE, got "hello"`,
		assert.Failure(t, e),
	)
}

func TestReplaceRefusesMissingDivider(t *testing.T) {
	path := goreplace_tester.Target(t, "alpha\n")
	e := goreplace.Run(
		path,
		"<<<<<<< SEARCH\nalpha\n>>>>>>> REPLACE\n",
		false,
		&bytes.Buffer{},
	)
	assert.StringContains(
		t,
		"block 1: no ======= before the end",
		assert.Failure(t, e),
	)
}

func TestReplaceKeepsFileMode(t *testing.T) {
	path := goreplace_tester.Target(t, "alpha\n")
	assert.FatalOnError(t, os.Chmod(path, 0o755))
	assert.FatalOnError(
		t,
		goreplace.Run(
			path,
			goreplace_tester.Blocks("alpha", "beta"),
			false,
			&bytes.Buffer{},
		),
	)
	i, e := os.Stat(path)
	assert.FatalOnError(t, e)
	assert.String(t, "-rwxr-xr-x", i.Mode().String())
}

func TestSpanUntilDeletesSectionAndKeepsEnd(t *testing.T) {
	path := goreplace_tester.Target(t, "# a\n\n## Alfa\nx\ny\n\n## Bravo\nz\n")
	var out bytes.Buffer
	assert.FatalOnError(
		t,
		goreplace.Run(
			path,
			goreplace_tester.Span(
				"## Alfa",
				textConstant.UntilMarker,
				"## Bravo",
				"",
			),
			false,
			&out,
		),
	)
	assert.String(t, "# a\n\n## Bravo\nz\n", system.ReadFileUnsafe(path))
	assert.StringContains(
		t,
		"@@ block 1, line 3\n-## Alfa\n-x\n-y\n-\n",
		out.String(),
	)
}

func TestSpanUntilKeepsEndOnItsOwnLine(t *testing.T) {
	path := goreplace_tester.Target(t, "## Alfa\nold\n## Bravo\n")
	assert.FatalOnError(
		t,
		goreplace.Run(
			path,
			goreplace_tester.Span(
				"## Alfa",
				textConstant.UntilMarker,
				"## Bravo",
				"## Alfa\nnew",
			),
			false,
			&bytes.Buffer{},
		),
	)
	assert.String(t, "## Alfa\nnew\n## Bravo\n", system.ReadFileUnsafe(path))
}

func TestSpanThroughRemovesEnd(t *testing.T) {
	path := goreplace_tester.Target(t, "a charlie b. delta. tail\n")
	assert.FatalOnError(
		t,
		goreplace.Run(
			path,
			goreplace_tester.Span(
				"charlie",
				textConstant.ThroughMarker,
				"delta.",
				"X",
			),
			false,
			&bytes.Buffer{},
		),
	)
	assert.String(t, "a X tail\n", system.ReadFileUnsafe(path))
}

func TestSpanToEndReplacesRest(t *testing.T) {
	path := goreplace_tester.Target(t, "keep\n## Echo\nold\n")
	assert.FatalOnError(
		t,
		goreplace.Run(
			path,
			goreplace_tester.Span(
				"## Echo",
				textConstant.ToEndMarker,
				"",
				"## Echo\nnone",
			),
			false,
			&bytes.Buffer{},
		),
	)
	assert.String(t, "keep\n## Echo\nnone\n", system.ReadFileUnsafe(path))
}

func TestSpanTakesFirstEndAfterStart(t *testing.T) {
	path := goreplace_tester.Target(
		t,
		"## End\n## Start\nx\n## End\ny\n## End\n",
	)
	assert.FatalOnError(
		t,
		goreplace.Run(
			path,
			goreplace_tester.Span(
				"## Start",
				textConstant.UntilMarker,
				"## End",
				"",
			),
			false,
			&bytes.Buffer{},
		),
	)
	assert.String(t, "## End\n## End\ny\n## End\n", system.ReadFileUnsafe(path))
}

func TestSpanRefusesMissingEnd(t *testing.T) {
	path := goreplace_tester.Target(t, "## Bravo\n## Alfa\nx\n")
	e := goreplace.Run(
		path,
		goreplace_tester.Span(
			"## Alfa",
			textConstant.UntilMarker,
			"## Bravo",
			"",
		),
		false,
		&bytes.Buffer{},
	)
	assert.StringContains(
		t,
		"block 1: end not found after the start at line 2",
		assert.Failure(t, e),
	)
	assert.String(t, "## Bravo\n## Alfa\nx\n", system.ReadFileUnsafe(path))
}

func TestSpanRefusesMissingClosing(t *testing.T) {
	path := goreplace_tester.Target(t, "alpha\n")
	e := goreplace.Run(
		path,
		"<<<<<<< SPAN\nalpha\n=======\n>>>>>>> REPLACE\n",
		false,
		&bytes.Buffer{},
	)
	assert.StringContains(
		t,
		"block 1: no ------- UNTIL, THROUGH or TO END",
		assert.Failure(t, e),
	)
}

func TestSpanRefusesTextAfterToEnd(t *testing.T) {
	path := goreplace_tester.Target(t, "alpha\n")
	e := goreplace.Run(
		path,
		goreplace_tester.Span("alpha", textConstant.ToEndMarker, "beta", ""),
		false,
		&bytes.Buffer{},
	)
	assert.StringContains(
		t,
		"block 1: nothing may follow ------- TO END",
		assert.Failure(t, e),
	)
}

func TestAfterInsertsBelowAnchorLine(t *testing.T) {
	path := goreplace_tester.Target(t, "| a | 1 |\n| b | 2 |\n")
	var out bytes.Buffer
	assert.FatalOnError(
		t,
		goreplace.Run(
			path,
			goreplace_tester.Insert(
				textConstant.AfterMarker,
				"| a |",
				"| c | 3 |",
			),
			false,
			&out,
		),
	)
	assert.String(
		t,
		"| a | 1 |\n| c | 3 |\n| b | 2 |\n",
		system.ReadFileUnsafe(path),
	)
	assert.StringContains(t, "@@ block 1, line 1\n+| c | 3 |\n", out.String())
}

func TestAfterLastLineWithoutNewline(t *testing.T) {
	path := goreplace_tester.Target(t, "alpha\nbeta")
	assert.FatalOnError(
		t,
		goreplace.Run(
			path,
			goreplace_tester.Insert(textConstant.AfterMarker, "beta", "gamma"),
			false,
			&bytes.Buffer{},
		),
	)
	assert.String(t, "alpha\nbeta\ngamma", system.ReadFileUnsafe(path))
}

func TestBeforeInsertsAboveAnchorLine(t *testing.T) {
	path := goreplace_tester.Target(t, "alpha\n| **Foxtrot** | | |\n")
	assert.FatalOnError(
		t,
		goreplace.Run(
			path,
			goreplace_tester.Insert(
				textConstant.BeforeMarker,
				"**Foxtrot**",
				"beta",
			),
			false,
			&bytes.Buffer{},
		),
	)
	assert.String(
		t,
		"alpha\nbeta\n| **Foxtrot** | | |\n",
		system.ReadFileUnsafe(path),
	)
}

func TestInsertAtSpanStartIsNoOverlap(t *testing.T) {
	path := goreplace_tester.Target(t, "a\n## S\nx\n## E\n")
	assert.FatalOnError(
		t,
		goreplace.Run(
			path,
			join.Empty(
				goreplace_tester.Span(
					"## S",
					textConstant.UntilMarker,
					"## E",
					"## T",
				),
				goreplace_tester.Insert(textConstant.AfterMarker, "a", "b"),
			),
			false,
			&bytes.Buffer{},
		),
	)
	assert.String(t, "a\nb\n## T\n## E\n", system.ReadFileUnsafe(path))
}

func TestInsertRefusesAnchorSpanRemoves(t *testing.T) {
	path := goreplace_tester.Target(t, "## S\nx\n## E\n")
	e := goreplace.Run(
		path,
		join.Empty(
			goreplace_tester.Span("## S", textConstant.UntilMarker, "## E", ""),
			goreplace_tester.Insert(textConstant.AfterMarker, "x", "y"),
		),
		false,
		&bytes.Buffer{},
	)
	assert.StringContains(
		t,
		"block 2: its anchor at line 2 is inside block 1",
		assert.Failure(t, e),
	)
	assert.String(t, "## S\nx\n## E\n", system.ReadFileUnsafe(path))
}

func TestAfterRefusesMultilineAnchor(t *testing.T) {
	path := goreplace_tester.Target(t, "alpha\nbeta\n")
	e := goreplace.Run(
		path,
		goreplace_tester.Insert(textConstant.AfterMarker, "alpha\nbeta", "x"),
		false,
		&bytes.Buffer{},
	)
	assert.StringContains(
		t,
		"block 1: an AFTER or BEFORE anchor is one line",
		assert.Failure(t, e),
	)
}

func TestAfterRefusesAmbiguousAnchor(t *testing.T) {
	path := goreplace_tester.Target(t, "| a |\n| a |\n")
	e := goreplace.Run(
		path,
		goreplace_tester.Insert(textConstant.AfterMarker, "| a |", "x"),
		false,
		&bytes.Buffer{},
	)
	assert.StringContains(
		t,
		"matches 2 times, at lines 1, 2",
		assert.Failure(t, e),
	)
}
