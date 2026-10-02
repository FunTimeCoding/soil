package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goflow/constant"
	"strings"
	"testing"
)

func TestAnOverrunningParagraphRewraps(t *testing.T) {
	assert.String(
		t,
		"one two three four five six seven eight\nnine ten\n",
		reflow(t, "one two three four five six seven eight nine ten\n"),
	)
}

func TestAParagraphInsideTheWidthIsUntouched(t *testing.T) {
	content := "one two\nthree four\n\nfive six\n"
	assert.String(t, content, reflow(t, content))
}

func TestRewrappingTwiceChangesNothingTheSecondTime(t *testing.T) {
	once := reflow(t, "one two three four five six seven eight nine ten\n")
	assert.String(t, once, reflow(t, once))
}

func TestAFencedBlockKeepsItsLongLines(t *testing.T) {
	content := "```\none two three four five six seven eight nine\n```\n"
	assert.String(t, content, reflow(t, content))
}

func TestATildeFenceKeepsItsLongLines(t *testing.T) {
	content := "~~~\none two three four five six seven eight nine\n~~~\n"
	assert.String(t, content, reflow(t, content))
}

func TestATableIsLeftAlone(t *testing.T) {
	content := "| one | two three four five six seven eight |\n|---|---|\n"
	assert.String(t, content, reflow(t, content))
}

func TestAPipeLineIsLeftAloneWithoutADelimiterRow(t *testing.T) {
	content := "| one | two three four five six seven eight |\n"
	assert.String(t, content, reflow(t, content))
}

func TestABlockquoteIsLeftAlone(t *testing.T) {
	content := "> one two three four five six seven eight nine ten\n"
	assert.String(t, content, reflow(t, content))
}

func TestAListItemRewrapsToItsOwnIndent(t *testing.T) {
	assert.String(
		t,
		"- one two three four five six seven\n  eight nine ten\n",
		reflow(t, "- one two three four five six seven eight nine ten\n"),
	)
}

func TestAnOrderedListItemRewrapsToItsOwnIndent(t *testing.T) {
	assert.String(
		t,
		"1. one two three four five six seven\n   eight nine ten\n",
		reflow(t, "1. one two three four five six seven eight nine ten\n"),
	)
}

func TestANestedListItemRewrapsToItsDeeperIndent(t *testing.T) {
	assert.String(
		t,
		"- outer\n  - inner one two three four five six\n    seven eight\n",
		reflow(
			t,
			"- outer\n  - inner one two three four five six seven eight\n",
		),
	)
}

func TestASecondParagraphInsideAListItemKeepsItsIndent(t *testing.T) {
	assert.String(
		t,
		"- one two\n\n  second paragraph one two three four\n  five six\n",
		reflow(
			t,
			"- one two\n\n  second paragraph one two three four five six\n",
		),
	)
}

func TestFrontMatterIsLeftAlone(t *testing.T) {
	content := "---\nbase: pkg/example and a long trailing phrase here\n---\n"
	assert.String(t, content, reflow(t, content))
}

func TestAnIndentedBlockIsLeftAlone(t *testing.T) {
	content := "    one two three four five six seven eight nine\n"
	assert.String(t, content, reflow(t, content))
}

func TestAHeadingIsLeftAlone(t *testing.T) {
	content := "# one two three four five six seven eight nine ten\n"
	assert.String(t, content, reflow(t, content))
}

func TestAWordLongerThanTheWidthIsNotBroken(t *testing.T) {
	assert.String(
		t,
		"short\ndoc/ai/design/example/a-path-longer-than-the-width.md\nafter\n",
		reflow(
			t,
			"short doc/ai/design/example/a-path-longer-than-the-width.md after\n",
		),
	)
}

func TestAParagraphJoinsBeforeItWraps(t *testing.T) {
	assert.String(
		t,
		"one two three four five six seven eight\nnine ten eleven\n",
		reflow(t, "one\ntwo three four five six seven eight nine ten eleven\n"),
	)
}

func TestBlankLinesSeparateParagraphs(t *testing.T) {
	assert.String(
		t,
		"one two three four five six seven eight\nnine\n\nten eleven\n",
		reflow(
			t,
			"one two three four five six seven eight nine\n\nten eleven\n",
		),
	)
}

func TestAFileWithoutATrailingNewlineKeepsItThatWay(t *testing.T) {
	assert.String(
		t,
		"one two three four five six seven eight\nnine",
		reflow(t, "one two three four five six seven eight nine"),
	)
}

func TestAnEmptyFileIsUntouched(t *testing.T) {
	assert.String(t, "", reflow(t, ""))
}

func TestAParagraphOpeningInBoldRewraps(t *testing.T) {
	assert.String(
		t,
		"**one** two three four five six seven\neight nine\n",
		reflow(t, "**one** two three four five six seven eight nine\n"),
	)
}

func TestAnAsteriskBulletRewrapsToItsOwnIndent(t *testing.T) {
	assert.String(
		t,
		"* one two three four five six seven\n  eight nine ten\n",
		reflow(t, "* one two three four five six seven eight nine ten\n"),
	)
}

func TestAThematicBreakIsNotABullet(t *testing.T) {
	content := "---\n\none two three four five six seven\n"
	assert.String(t, content, reflow(t, content))
}

func TestAnOpeningDelimiterWithNoCloseIsNotFrontMatter(t *testing.T) {
	assert.String(
		t,
		"---\none two three four five six seven eight\nnine\n",
		reflow(t, "---\none two three four five six seven eight nine\n"),
	)
}

func TestEveryWordSurvivesAReflow(t *testing.T) {
	content := join.NewLine(
		[]string{
			constant.Delimiter,
			"base: pkg/one",
			constant.Delimiter,
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
		strings.Fields(reflow(t, content)),
	)
}
