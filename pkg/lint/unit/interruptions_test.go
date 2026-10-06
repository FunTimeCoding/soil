package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	lintConstant "github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/reflow"
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
	"testing"
)

func TestAProseDashUnderAParagraphLineIsFlagged(t *testing.T) {
	assert.Any(t, []int{2}, reflow.Interruptions("one two\n- three four\n"))
}

func TestAProsePlusUnderAParagraphLineIsFlagged(t *testing.T) {
	assert.Any(t, []int{2}, reflow.Interruptions("one two\n+ three four\n"))
}

func TestANestedProseDashIsFlagged(t *testing.T) {
	assert.Any(t, []int{2}, reflow.Interruptions("- one two\n  - three four\n"))
}

func TestAListAfterABlankLineIsClean(t *testing.T) {
	assert.Any(t, []int(nil), reflow.Interruptions("one two\n\n- three\n"))
}

func TestAListAfterAColonIsClean(t *testing.T) {
	assert.Any(t, []int(nil), reflow.Interruptions("Items:\n- one\n- two\n"))
}

func TestAListAfterABoldColonIsClean(t *testing.T) {
	assert.Any(t, []int(nil), reflow.Interruptions("**Items:**\n- one\n"))
}

func TestAQuotedProseDashIsFlagged(t *testing.T) {
	assert.Any(t, []int{2}, reflow.Interruptions("one two\n- \"three\" four\n"))
}

func TestAParenthesisedProseDashIsFlagged(t *testing.T) {
	assert.Any(t, []int{2}, reflow.Interruptions("one two\n- (three) four\n"))
}

func TestAQuotedCapitalisedItemIsClean(t *testing.T) {
	assert.Any(t, []int(nil), reflow.Interruptions("one two\n- \"Three\"\n"))
}

func TestACodeItemIsClean(t *testing.T) {
	assert.Any(t, []int(nil), reflow.Interruptions("one two\n- `three`\n"))
}

func TestACapitalisedItemIsClean(t *testing.T) {
	assert.Any(t, []int(nil), reflow.Interruptions("one two\n- Three\n"))
}

func TestFrontMatterIsNeverFlagged(t *testing.T) {
	assert.Any(
		t,
		[]int(nil),
		reflow.Interruptions("---\nkey: value\n- item\n---\n\nbody\n"),
	)
}

func TestAnInterruptingListIsAConcernThatIsNotFixed(t *testing.T) {
	l := lint.InterruptingList(
		constant.UpperAlfa,
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
				Type:     lintConstant.ConcernLine,
				Line:     2,
				LineText: "- three four",
				Fixed:    false,
			},
		},
		"",
		l,
	)
}
