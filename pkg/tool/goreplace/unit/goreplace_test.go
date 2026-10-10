package unit

import (
	"bytes"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goreplace"
	"github.com/funtimecoding/soil/pkg/tool/goreplace/constant"
	"os"
	"testing"
)

func TestReplaceAppliesEveryBlock(t *testing.T) {
	path := target(t, "alpha\nbeta\ngamma\n")
	var out bytes.Buffer
	assert.FatalOnError(
		t,
		goreplace.Run(
			path,
			blocks("alpha", "ALPHA", "gamma", "GAMMA"),
			false,
			&out,
		),
	)
	assert.String(t, "ALPHA\nbeta\nGAMMA\n", content(t, path))
	assert.StringContains(t, "@@ block 2, line 3\n-gamma\n+GAMMA", out.String())
	assert.StringContains(t, "2 blocks applied", out.String())
}

func TestReplaceDryRunWritesNothing(t *testing.T) {
	path := target(t, "alpha\nbeta\n")
	var out bytes.Buffer
	assert.FatalOnError(
		t,
		goreplace.Run(path, blocks("alpha", "ALPHA"), true, &out),
	)
	assert.String(t, "alpha\nbeta\n", content(t, path))
	assert.StringContains(t, "would apply (dry run", out.String())
}

func TestReplaceWritesNothingWhenOneBlockFails(t *testing.T) {
	path := target(t, "alpha\nbeta\n")
	e := goreplace.Run(
		path,
		blocks("alpha", "ALPHA", "missing", "x"),
		false,
		&bytes.Buffer{},
	)
	assert.StringContains(t, "block 2: search not found", failure(t, e))
	assert.String(t, "alpha\nbeta\n", content(t, path))
}

func TestReplaceRefusesAnAmbiguousSearch(t *testing.T) {
	path := target(t, "x\ny\nx\n")
	e := goreplace.Run(path, blocks("x", "z"), false, &bytes.Buffer{})
	assert.StringContains(t, "matches 2 times, at lines 1, 3", failure(t, e))
}

func TestReplaceRefusesOverlappingBlocks(t *testing.T) {
	path := target(t, "alpha\nbeta\ngamma\n")
	e := goreplace.Run(
		path,
		blocks("alpha\nbeta", "x", "beta\ngamma", "y"),
		false,
		&bytes.Buffer{},
	)
	assert.StringContains(t, "blocks 1 and 2 overlap at line 2", failure(t, e))
}

func TestReplaceMatchesAgainstTheOriginal(t *testing.T) {
	path := target(t, "alpha\nbeta\n")
	e := goreplace.Run(
		path,
		blocks("alpha", "delta", "delta", "epsilon"),
		false,
		&bytes.Buffer{},
	)
	assert.StringContains(t, "block 2: search not found", failure(t, e))
}

func TestReplaceDeletesAWholeLine(t *testing.T) {
	path := target(t, "alpha\nbeta\ngamma\n")
	assert.FatalOnError(
		t,
		goreplace.Run(path, blocks("beta\n", ""), false, &bytes.Buffer{}),
	)
	assert.String(t, "alpha\ngamma\n", content(t, path))
}

func TestReplaceHintsAtWhitespace(t *testing.T) {
	path := target(t, "  indented line\n")
	e := goreplace.Run(
		path,
		blocks("indented  line", "x"),
		false,
		&bytes.Buffer{},
	)
	assert.StringContains(
		t,
		"matches when whitespace is ignored",
		failure(t, e),
	)
}

func TestReplaceHintsAtTheFirstLine(t *testing.T) {
	path := target(t, "one\ntwo\n")
	e := goreplace.Run(path, blocks("one\nthree", "x"), false, &bytes.Buffer{})
	assert.StringContains(t, "its first line occurs at line 1", failure(t, e))
}

func TestReplaceRefusesStrayText(t *testing.T) {
	path := target(t, "alpha\n")
	e := goreplace.Run(path, "hello\n", false, &bytes.Buffer{})
	assert.StringContains(
		t,
		`line 1: expected <<<<<<< SEARCH, SPAN, AFTER or BEFORE, got "hello"`,
		failure(t, e),
	)
}

func TestReplaceRefusesAMissingDivider(t *testing.T) {
	path := target(t, "alpha\n")
	e := goreplace.Run(
		path,
		"<<<<<<< SEARCH\nalpha\n>>>>>>> REPLACE\n",
		false,
		&bytes.Buffer{},
	)
	assert.StringContains(
		t,
		"block 1: no ======= before the end",
		failure(t, e),
	)
}

func TestReplaceKeepsTheFileMode(t *testing.T) {
	path := target(t, "alpha\n")
	assert.FatalOnError(t, os.Chmod(path, 0o755))
	assert.FatalOnError(
		t,
		goreplace.Run(path, blocks("alpha", "beta"), false, &bytes.Buffer{}),
	)
	i, e := os.Stat(path)
	assert.FatalOnError(t, e)
	assert.String(t, "-rwxr-xr-x", i.Mode().String())
}

func TestSpanUntilDeletesASectionAndKeepsTheEnd(t *testing.T) {
	path := target(t, "# a\n\n## Alfa\nx\ny\n\n## Bravo\nz\n")
	var out bytes.Buffer
	assert.FatalOnError(
		t,
		goreplace.Run(
			path,
			spanBlock("## Alfa", constant.UntilMarker, "## Bravo", ""),
			false,
			&out,
		),
	)
	assert.String(t, "# a\n\n## Bravo\nz\n", content(t, path))
	assert.StringContains(
		t,
		"@@ block 1, line 3\n-## Alfa\n-x\n-y\n-\n",
		out.String(),
	)
}

func TestSpanUntilKeepsTheEndOnItsOwnLine(t *testing.T) {
	path := target(t, "## Alfa\nold\n## Bravo\n")
	assert.FatalOnError(
		t,
		goreplace.Run(
			path,
			spanBlock(
				"## Alfa",
				constant.UntilMarker,
				"## Bravo",
				"## Alfa\nnew",
			),
			false,
			&bytes.Buffer{},
		),
	)
	assert.String(t, "## Alfa\nnew\n## Bravo\n", content(t, path))
}

func TestSpanThroughRemovesTheEnd(t *testing.T) {
	path := target(t, "a charlie b. delta. tail\n")
	assert.FatalOnError(
		t,
		goreplace.Run(
			path,
			spanBlock("charlie", constant.ThroughMarker, "delta.", "X"),
			false,
			&bytes.Buffer{},
		),
	)
	assert.String(t, "a X tail\n", content(t, path))
}

func TestSpanToEndReplacesTheRest(t *testing.T) {
	path := target(t, "keep\n## Echo\nold\n")
	assert.FatalOnError(
		t,
		goreplace.Run(
			path,
			spanBlock("## Echo", constant.ToEndMarker, "", "## Echo\nnone"),
			false,
			&bytes.Buffer{},
		),
	)
	assert.String(t, "keep\n## Echo\nnone\n", content(t, path))
}

func TestSpanTakesTheFirstEndAfterTheStart(t *testing.T) {
	path := target(t, "## End\n## Start\nx\n## End\ny\n## End\n")
	assert.FatalOnError(
		t,
		goreplace.Run(
			path,
			spanBlock("## Start", constant.UntilMarker, "## End", ""),
			false,
			&bytes.Buffer{},
		),
	)
	assert.String(t, "## End\n## End\ny\n## End\n", content(t, path))
}

func TestSpanRefusesAMissingEnd(t *testing.T) {
	path := target(t, "## Bravo\n## Alfa\nx\n")
	e := goreplace.Run(
		path,
		spanBlock("## Alfa", constant.UntilMarker, "## Bravo", ""),
		false,
		&bytes.Buffer{},
	)
	assert.StringContains(
		t,
		"block 1: end not found after the start at line 2",
		failure(t, e),
	)
	assert.String(t, "## Bravo\n## Alfa\nx\n", content(t, path))
}

func TestSpanRefusesAMissingClosing(t *testing.T) {
	path := target(t, "alpha\n")
	e := goreplace.Run(
		path,
		"<<<<<<< SPAN\nalpha\n=======\n>>>>>>> REPLACE\n",
		false,
		&bytes.Buffer{},
	)
	assert.StringContains(
		t,
		"block 1: no ------- UNTIL, THROUGH or TO END",
		failure(t, e),
	)
}

func TestSpanRefusesTextAfterToEnd(t *testing.T) {
	path := target(t, "alpha\n")
	e := goreplace.Run(
		path,
		spanBlock("alpha", constant.ToEndMarker, "beta", ""),
		false,
		&bytes.Buffer{},
	)
	assert.StringContains(
		t,
		"block 1: nothing may follow ------- TO END",
		failure(t, e),
	)
}

func TestAfterInsertsBelowTheAnchorLine(t *testing.T) {
	path := target(t, "| a | 1 |\n| b | 2 |\n")
	var out bytes.Buffer
	assert.FatalOnError(
		t,
		goreplace.Run(
			path,
			insertBlock(constant.AfterMarker, "| a |", "| c | 3 |"),
			false,
			&out,
		),
	)
	assert.String(t, "| a | 1 |\n| c | 3 |\n| b | 2 |\n", content(t, path))
	assert.StringContains(t, "@@ block 1, line 1\n+| c | 3 |\n", out.String())
}

func TestAfterTheLastLineWithoutANewline(t *testing.T) {
	path := target(t, "alpha\nbeta")
	assert.FatalOnError(
		t,
		goreplace.Run(
			path,
			insertBlock(constant.AfterMarker, "beta", "gamma"),
			false,
			&bytes.Buffer{},
		),
	)
	assert.String(t, "alpha\nbeta\ngamma", content(t, path))
}

func TestBeforeInsertsAboveTheAnchorLine(t *testing.T) {
	path := target(t, "alpha\n| **Foxtrot** | | |\n")
	assert.FatalOnError(
		t,
		goreplace.Run(
			path,
			insertBlock(constant.BeforeMarker, "**Foxtrot**", "beta"),
			false,
			&bytes.Buffer{},
		),
	)
	assert.String(t, "alpha\nbeta\n| **Foxtrot** | | |\n", content(t, path))
}

func TestInsertAtASpanStartIsNoOverlap(t *testing.T) {
	path := target(t, "a\n## S\nx\n## E\n")
	assert.FatalOnError(
		t,
		goreplace.Run(
			path,
			join.Empty(
				spanBlock("## S", constant.UntilMarker, "## E", "## T"),
				insertBlock(constant.AfterMarker, "a", "b"),
			),
			false,
			&bytes.Buffer{},
		),
	)
	assert.String(t, "a\nb\n## T\n## E\n", content(t, path))
}

func TestInsertRefusesAnAnchorASpanRemoves(t *testing.T) {
	path := target(t, "## S\nx\n## E\n")
	e := goreplace.Run(
		path,
		join.Empty(
			spanBlock("## S", constant.UntilMarker, "## E", ""),
			insertBlock(constant.AfterMarker, "x", "y"),
		),
		false,
		&bytes.Buffer{},
	)
	assert.StringContains(
		t,
		"block 2: its anchor at line 2 is inside block 1",
		failure(t, e),
	)
	assert.String(t, "## S\nx\n## E\n", content(t, path))
}

func TestAfterRefusesAMultilineAnchor(t *testing.T) {
	path := target(t, "alpha\nbeta\n")
	e := goreplace.Run(
		path,
		insertBlock(constant.AfterMarker, "alpha\nbeta", "x"),
		false,
		&bytes.Buffer{},
	)
	assert.StringContains(
		t,
		"block 1: an AFTER or BEFORE anchor is one line",
		failure(t, e),
	)
}

func TestAfterRefusesAnAmbiguousAnchor(t *testing.T) {
	path := target(t, "| a |\n| a |\n")
	e := goreplace.Run(
		path,
		insertBlock(constant.AfterMarker, "| a |", "x"),
		false,
		&bytes.Buffer{},
	)
	assert.StringContains(t, "matches 2 times, at lines 1, 2", failure(t, e))
}
