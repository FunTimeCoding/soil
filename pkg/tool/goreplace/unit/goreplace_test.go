package unit

import (
	"bytes"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goreplace"
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
		`line 1: expected <<<<<<< SEARCH, got "hello"`,
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
