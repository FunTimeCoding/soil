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

func TestAnOverrunningParagraphRewraps(t *testing.T) {
	assert.String(
		t,
		"one two three four five six seven eight\nnine ten\n",
		rewrapped(t, "one two three four five six seven eight nine ten\n"),
	)
}

func TestAParagraphInsideTheWidthIsUntouched(t *testing.T) {
	content := "one two\nthree four\n\nfive six\n"
	assert.String(t, content, rewrapped(t, content))
}

func TestRewrappingTwiceChangesNothingTheSecondTime(t *testing.T) {
	once := rewrapped(t, "one two three four five six seven eight nine ten\n")
	assert.String(t, once, rewrapped(t, once))
}

func TestAFencedBlockKeepsItsLongLines(t *testing.T) {
	content := "```\none two three four five six seven eight nine\n```\n"
	assert.String(t, content, rewrapped(t, content))
}

func TestATildeFenceKeepsItsLongLines(t *testing.T) {
	content := "~~~\none two three four five six seven eight nine\n~~~\n"
	assert.String(t, content, rewrapped(t, content))
}

func TestATableIsLeftAlone(t *testing.T) {
	content := "| one | two three four five six seven eight |\n|---|---|\n"
	assert.String(t, content, rewrapped(t, content))
}

func TestAPipeLineIsLeftAloneWithoutADelimiterRow(t *testing.T) {
	content := "| one | two three four five six seven eight |\n"
	assert.String(t, content, rewrapped(t, content))
}

func TestABlockquoteIsLeftAlone(t *testing.T) {
	content := "> one two three four five six seven eight nine ten\n"
	assert.String(t, content, rewrapped(t, content))
}

func TestAListItemRewrapsToItsOwnIndent(t *testing.T) {
	assert.String(
		t,
		"- one two three four five six seven\n  eight nine ten\n",
		rewrapped(t, "- one two three four five six seven eight nine ten\n"),
	)
}

func TestAnOrderedListItemRewrapsToItsOwnIndent(t *testing.T) {
	assert.String(
		t,
		"1. one two three four five six seven\n   eight nine ten\n",
		rewrapped(t, "1. one two three four five six seven eight nine ten\n"),
	)
}

func TestANestedListItemRewrapsToItsDeeperIndent(t *testing.T) {
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

func TestASecondParagraphInsideAListItemKeepsItsIndent(t *testing.T) {
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

func TestAnIndentedBlockIsLeftAlone(t *testing.T) {
	content := "    one two three four five six seven eight nine\n"
	assert.String(t, content, rewrapped(t, content))
}

func TestAHeadingIsLeftAlone(t *testing.T) {
	content := "# one two three four five six seven eight nine ten\n"
	assert.String(t, content, rewrapped(t, content))
}

func TestAWordLongerThanTheWidthIsNotBroken(t *testing.T) {
	assert.String(
		t,
		"short\ndoc/ai/design/example/a-path-longer-than-the-width.md\nafter\n",
		rewrapped(
			t,
			"short doc/ai/design/example/a-path-longer-than-the-width.md after\n",
		),
	)
}

func TestAParagraphJoinsBeforeItWraps(t *testing.T) {
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

func TestAFileWithoutATrailingNewlineKeepsItThatWay(t *testing.T) {
	assert.String(
		t,
		"one two three four five six seven eight\nnine",
		rewrapped(t, "one two three four five six seven eight nine"),
	)
}

func TestAnEmptyFileIsUntouched(t *testing.T) {
	assert.String(t, "", rewrapped(t, ""))
}

func TestAParagraphOpeningInBoldRewraps(t *testing.T) {
	assert.String(
		t,
		"**one** two three four five six seven\neight nine\n",
		rewrapped(t, "**one** two three four five six seven eight nine\n"),
	)
}

func TestAnAsteriskBulletRewrapsToItsOwnIndent(t *testing.T) {
	assert.String(
		t,
		"* one two three four five six seven\n  eight nine ten\n",
		rewrapped(t, "* one two three four five six seven eight nine ten\n"),
	)
}

func TestAThematicBreakIsNotABullet(t *testing.T) {
	content := "---\n\none two three four five six seven\n"
	assert.String(t, content, rewrapped(t, content))
}

func TestAnOpeningDelimiterWithNoCloseIsNotFrontMatter(t *testing.T) {
	assert.String(
		t,
		"---\none two three four five six seven eight\nnine\n",
		rewrapped(t, "---\none two three four five six seven eight nine\n"),
	)
}

func TestWhitespaceInsideACodeSpanSurvives(t *testing.T) {
	assert.String(
		t,
		"one two three four `a  b` five six seven\neight\n",
		rewrapped(t, "one two three four `a  b` five six seven eight\n"),
	)
}

func TestACodeSpanIsNeverSplitAcrossLines(t *testing.T) {
	assert.String(
		t,
		"one two three four five six\n`go run main.go` seven\n",
		rewrapped(t, "one two three four five six `go run main.go` seven\n"),
	)
}

func TestACodeSpanAlreadySplitIsRejoined(t *testing.T) {
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

func TestADashTravelsWithTheWordBeforeIt(t *testing.T) {
	assert.String(
		t,
		"one two three four five six seven\neight - nine\n",
		rewrapped(t, "one two three four five six seven eight - nine\n"),
	)
}

func TestAnOrderedMarkerTravelsWithTheWordBeforeIt(t *testing.T) {
	assert.String(
		t,
		"one two three four five six seven\neight 1. nine\n",
		rewrapped(t, "one two three four five six seven eight 1. nine\n"),
	)
}

func TestAHashTravelsWithTheWordBeforeIt(t *testing.T) {
	assert.String(
		t,
		"one two three four five six seven\neight ## nine\n",
		rewrapped(t, "one two three four five six seven eight ## nine\n"),
	)
}

func TestAQuoteMarkTravelsWithTheWordBeforeIt(t *testing.T) {
	assert.String(
		t,
		"one two three four five six seven\neight >nine\n",
		rewrapped(t, "one two three four five six seven eight >nine\n"),
	)
}

func TestADecimalAtALineStartStaysProse(t *testing.T) {
	assert.String(
		t,
		"one two three four five six seven eight\n2.5 nine\n",
		rewrapped(t, "one two three four five six seven eight 2.5 nine\n"),
	)
}

func TestAWrapThatWouldDropAHardLineBreakIsRefused(t *testing.T) {
	_, e := reflow.Reflow(
		"one two three four five six seven eight  \nnine\n",
		constant.FixtureReflowWidth,
	)
	assert.Error(t, e)
}

func TestEveryWordSurvivesAReflow(t *testing.T) {
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
