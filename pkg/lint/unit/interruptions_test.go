package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/reflow"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
	"testing"
)

func TestProseDashUnderParagraphLineIsFlagged(t *testing.T) {
	assert.Any(t, []int{2}, reflow.Interruptions("one two\n- three four\n"))
}

func TestProsePlusUnderParagraphLineIsFlagged(t *testing.T) {
	assert.Any(t, []int{2}, reflow.Interruptions("one two\n+ three four\n"))
}

func TestNestedProseDashIsFlagged(t *testing.T) {
	assert.Any(t, []int{2}, reflow.Interruptions("- one two\n  - three four\n"))
}

func TestListAfterBlankLineIsClean(t *testing.T) {
	assert.Any(t, []int(nil), reflow.Interruptions("one two\n\n- three\n"))
}

func TestListAfterColonIsClean(t *testing.T) {
	assert.Any(t, []int(nil), reflow.Interruptions("Items:\n- one\n- two\n"))
}

func TestListAfterBoldColonIsClean(t *testing.T) {
	assert.Any(t, []int(nil), reflow.Interruptions("**Items:**\n- one\n"))
}

func TestQuotedProseDashIsFlagged(t *testing.T) {
	assert.Any(t, []int{2}, reflow.Interruptions("one two\n- \"three\" four\n"))
}

func TestParenthesisedProseDashIsFlagged(t *testing.T) {
	assert.Any(t, []int{2}, reflow.Interruptions("one two\n- (three) four\n"))
}

func TestQuotedCapitalisedItemIsClean(t *testing.T) {
	assert.Any(t, []int(nil), reflow.Interruptions("one two\n- \"Three\"\n"))
}

func TestCodeItemIsClean(t *testing.T) {
	assert.Any(t, []int(nil), reflow.Interruptions("one two\n- `three`\n"))
}

func TestCapitalisedItemIsClean(t *testing.T) {
	assert.Any(t, []int(nil), reflow.Interruptions("one two\n- Three\n"))
}

func TestFrontMatterIsNeverFlagged(t *testing.T) {
	assert.Any(
		t,
		[]int(nil),
		reflow.Interruptions("---\nkey: value\n- item\n---\n\nbody\n"),
	)
}

func TestInterruptingListIsConcernThatIsNotFixed(t *testing.T) {
	l := lint.InterruptingList(
		stringsConstant.UpperAlfa,
		strings.NewReader("one two\n- three four\n"),
	)
	assertReport(
		t,
		"Alfa",
		true,
		[]*concern.Concern{
			{
				Key:      "interrupting_list",
				Text:     "List starts directly under a paragraph line - join a prose dash back, or put a blank line or a colon before a real list",
				Path:     "Alfa",
				Type:     constant.ConcernLine,
				Line:     2,
				LineText: "- three four",
				Fixed:    false,
			},
		},
		"",
		l,
	)
}
