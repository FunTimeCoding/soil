package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/reflow"
	markupConstant "github.com/funtimecoding/soil/pkg/markup/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"strings"
	"testing"
)

func TestOverrunningParagraphRewraps(t *testing.T) {
	assert.String(
		t,
		"one two three four five six seven eight\nnine ten\n",
		rewrapped(t, "one two three four five six seven eight nine ten\n"),
	)
}

func TestParagraphInsideWidthIsUntouched(t *testing.T) {
	content := "one two\nthree four\n\nfive six\n"
	assert.String(t, content, rewrapped(t, content))
}

func TestRewrappingTwiceChangesNothingSecondTime(t *testing.T) {
	once := rewrapped(t, "one two three four five six seven eight nine ten\n")
	assert.String(t, once, rewrapped(t, once))
}

func TestFencedBlockKeepsItsLongLines(t *testing.T) {
	content := "```\none two three four five six seven eight nine\n```\n"
	assert.String(t, content, rewrapped(t, content))
}

func TestTildeFenceKeepsItsLongLines(t *testing.T) {
	content := "~~~\none two three four five six seven eight nine\n~~~\n"
	assert.String(t, content, rewrapped(t, content))
}

func TestTableIsLeftAlone(t *testing.T) {
	content := "| one | two three four five six seven eight |\n|---|---|\n"
	assert.String(t, content, rewrapped(t, content))
}

func TestPipeLineIsLeftAloneWithoutDelimiterRow(t *testing.T) {
	content := "| one | two three four five six seven eight |\n"
	assert.String(t, content, rewrapped(t, content))
}

func TestBlockquoteIsLeftAlone(t *testing.T) {
	content := "> one two three four five six seven eight nine ten\n"
	assert.String(t, content, rewrapped(t, content))
}

func TestListItemRewrapsToItsOwnIndent(t *testing.T) {
	assert.String(
		t,
		"- one two three four five six seven\n  eight nine ten\n",
		rewrapped(t, "- one two three four five six seven eight nine ten\n"),
	)
}

func TestOrderedListItemRewrapsToItsOwnIndent(t *testing.T) {
	assert.String(
		t,
		"1. one two three four five six seven\n   eight nine ten\n",
		rewrapped(t, "1. one two three four five six seven eight nine ten\n"),
	)
}

func TestNestedListItemRewrapsToItsDeeperIndent(t *testing.T) {
	assert.String(
		t,
		"- outer\n  - inner one two three four five six\n    seven eight\n",
		rewrapped(
			t,
			"- outer\n  - inner one two three four five six seven eight\n",
		),
	)
}

func TestNeighbouringItemsStayApartWhenOneRewraps(t *testing.T) {
	assert.String(
		t,
		"- **One:** short\n- **Two:** one two three four five six\n  seven eight\n- **Three:** short\n",
		rewrapped(
			t,
			"- **One:** short\n- **Two:** one two three four five six seven eight\n- **Three:** short\n",
		),
	)
}

func TestSecondParagraphInsideListItemKeepsItsIndent(t *testing.T) {
	assert.String(
		t,
		"- one two\n\n  second paragraph one two three four\n  five six\n",
		rewrapped(
			t,
			"- one two\n\n  second paragraph one two three four five six\n",
		),
	)
}

func TestFrontMatterIsLeftAlone(t *testing.T) {
	content := "---\nbase: pkg/example and a long trailing phrase here\n---\n"
	assert.String(t, content, rewrapped(t, content))
}

func TestIndentedBlockIsLeftAlone(t *testing.T) {
	content := "    one two three four five six seven eight nine\n"
	assert.String(t, content, rewrapped(t, content))
}

func TestHeadingIsLeftAlone(t *testing.T) {
	content := "# one two three four five six seven eight nine ten\n"
	assert.String(t, content, rewrapped(t, content))
}

func TestWordLongerThanWidthIsNotBroken(t *testing.T) {
	assert.String(
		t,
		"short\ndoc/ai/design/example/a-path-longer-than-the-width.md\nafter\n",
		rewrapped(
			t,
			"short doc/ai/design/example/a-path-longer-than-the-width.md after\n",
		),
	)
}

func TestParagraphJoinsBeforeItWraps(t *testing.T) {
	assert.String(
		t,
		"one two three four five six seven eight\nnine ten eleven\n",
		rewrapped(
			t,
			"one\ntwo three four five six seven eight nine ten eleven\n",
		),
	)
}

func TestBlankLinesSeparateParagraphs(t *testing.T) {
	assert.String(
		t,
		"one two three four five six seven eight\nnine\n\nten eleven\n",
		rewrapped(
			t,
			"one two three four five six seven eight nine\n\nten eleven\n",
		),
	)
}

func TestFileWithoutTrailingNewlineKeepsItThatWay(t *testing.T) {
	assert.String(
		t,
		"one two three four five six seven eight\nnine",
		rewrapped(t, "one two three four five six seven eight nine"),
	)
}

func TestEmptyFileIsUntouched(t *testing.T) {
	assert.String(t, "", rewrapped(t, ""))
}

func TestParagraphOpeningInBoldRewraps(t *testing.T) {
	assert.String(
		t,
		"**one** two three four five six seven\neight nine\n",
		rewrapped(t, "**one** two three four five six seven eight nine\n"),
	)
}

func TestAsteriskBulletRewrapsToItsOwnIndent(t *testing.T) {
	assert.String(
		t,
		"* one two three four five six seven\n  eight nine ten\n",
		rewrapped(t, "* one two three four five six seven eight nine ten\n"),
	)
}

func TestThematicBreakIsNotBullet(t *testing.T) {
	content := "---\n\none two three four five six seven\n"
	assert.String(t, content, rewrapped(t, content))
}

func TestOpeningDelimiterWithNoCloseIsNotFrontMatter(t *testing.T) {
	assert.String(
		t,
		"---\none two three four five six seven eight\nnine\n",
		rewrapped(t, "---\none two three four five six seven eight nine\n"),
	)
}

func TestWhitespaceInsideCodeSpanSurvives(t *testing.T) {
	assert.String(
		t,
		"one two three four `a  b` five six seven\neight\n",
		rewrapped(t, "one two three four `a  b` five six seven eight\n"),
	)
}

func TestCodeSpanIsNeverSplitAcrossLines(t *testing.T) {
	assert.String(
		t,
		"one two three four five six\n`go run main.go` seven\n",
		rewrapped(t, "one two three four five six `go run main.go` seven\n"),
	)
}

func TestLinkIsNeverSplitAcrossLines(t *testing.T) {
	assert.String(
		t,
		"one two three four five six\n[seven eight](https://e.x/a) nine\n",
		rewrapped(
			t,
			"one two three four five six [seven eight](https://e.x/a) nine\n",
		),
	)
}

func TestLinkLongerThanWidthStandsAlone(t *testing.T) {
	assert.String(
		t,
		"short\n[a link text](https://example.net/a/path/longer/than/forty)\nafter\n",
		rewrapped(
			t,
			"short [a link text](https://example.net/a/path/longer/than/forty) after\n",
		),
	)
}

func TestLinkAlreadySplitIsRejoined(t *testing.T) {
	assert.String(
		t,
		"one [seven eight](https://e.x/a) two\nthree four five six seven eight nine ten\n",
		rewrapped(
			t,
			"one [seven\neight](https://e.x/a) two three four five six seven eight nine ten\n",
		),
	)
}

func TestLinkWithTitleIsOneUnit(t *testing.T) {
	assert.String(
		t,
		"one two three four five six\n[seven](https://e.x/a \"t u\") nine\n",
		rewrapped(
			t,
			"one two three four five six [seven](https://e.x/a \"t u\") nine\n",
		),
	)
}

func TestImageIsNeverSplitAcrossLines(t *testing.T) {
	assert.String(
		t,
		"one two three four five six\n![seven eight](https://e.x/a) nine\n",
		rewrapped(
			t,
			"one two three four five six ![seven eight](https://e.x/a) nine\n",
		),
	)
}

func TestCodeSpanAlreadySplitIsRejoined(t *testing.T) {
	assert.String(
		t,
		"one `go run main.go` two three four five\nsix seven eight nine\n",
		rewrapped(
			t,
			"one `go run\nmain.go` two three four five six seven eight nine\n",
		),
	)
}

func TestWidthCountsCharactersNotBytes(t *testing.T) {
	content := "→ → → → → → → → → → → → → → → → → → → →\n"
	assert.String(t, content, rewrapped(t, content))
}

func TestDashTravelsWithWordBeforeIt(t *testing.T) {
	assert.String(
		t,
		"one two three four five six seven\neight - nine\n",
		rewrapped(t, "one two three four five six seven eight - nine\n"),
	)
}

func TestOrderedMarkerTravelsWithWordBeforeIt(t *testing.T) {
	assert.String(
		t,
		"one two three four five six seven\neight 1. nine\n",
		rewrapped(t, "one two three four five six seven eight 1. nine\n"),
	)
}

func TestHashTravelsWithWordBeforeIt(t *testing.T) {
	assert.String(
		t,
		"one two three four five six seven\neight ## nine\n",
		rewrapped(t, "one two three four five six seven eight ## nine\n"),
	)
}

func TestQuoteMarkTravelsWithWordBeforeIt(t *testing.T) {
	assert.String(
		t,
		"one two three four five six seven\neight >nine\n",
		rewrapped(t, "one two three four five six seven eight >nine\n"),
	)
}

func TestDecimalAtLineStartStaysProse(t *testing.T) {
	assert.String(
		t,
		"one two three four five six seven eight\n2.5 nine\n",
		rewrapped(t, "one two three four five six seven eight 2.5 nine\n"),
	)
}

func TestWrapThatWouldDropHardLineBreakIsRefused(t *testing.T) {
	_, e := reflow.Reflow(
		"one two three four five six seven eight  \nnine\n",
		constant.FixtureReflowWidth,
	)
	assert.Error(t, e)
}

func TestEveryWordSurvivesReflow(t *testing.T) {
	content := join.NewLine(
		[]string{
			markupConstant.FrontMatterDelimiter,
			"base: pkg/one",
			markupConstant.FrontMatterDelimiter,
			"",
			"# Heading one two three",
			"",
			"one two three four five six seven eight nine ten eleven",
			"",
			"| one | two three four five six seven |",
			"|---|---|",
			"",
			"```",
			"literal one two three four five six seven eight",
			"```",
			"",
			"- one two three four five six seven eight nine",
			"  ten eleven",
			"",
			"> quoted one two three four five six seven eight",
			"",
		},
	)
	assert.Strings(
		t,
		strings.Fields(content),
		strings.Fields(rewrapped(t, content)),
	)
}
