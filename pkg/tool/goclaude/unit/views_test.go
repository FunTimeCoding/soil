package unit

import (
	"errors"
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"strings"
	"testing"
	"unicode/utf8"
)

func TestAnUnknownSessionExitsNotFound(t *testing.T) {
	for _, arguments := range sessionCommands("zulu") {
		_, failure, code := run(t, newMock(t), arguments...)
		assert.Integer(t, 1, code)
		assert.StringContains(t, "zulu", failure)
		assert.StringContains(t, "not found", failure)
	}
}

func TestAnAmbiguousSessionListsTheMatches(t *testing.T) {
	c := newMock(t)
	c.Matches = []client.ResolveMatch{
		{Identifier: "alfa1234-0000", Name: new("kilo")},
		{Identifier: "alfa5678-0000", Alias: new("lima")},
	}
	_, failure, code := run(t, c, "session", "show", "alfa")
	assert.Integer(t, 1, code)
	assert.StringContains(t, `ambiguous: "alfa" matches 2 sessions`, failure)
	assert.StringContains(t, "alfa1234", failure)
	assert.StringContains(t, "lima", failure)
}

func TestAFailedCommandExitsWithItsMessage(t *testing.T) {
	for _, arguments := range append(
		sessionCommands("alfa"),
		[]string{"cost"},
		[]string{"status"},
		[]string{"usage"},
		[]string{"session", soil.List},
		[]string{"session", "heatmap"},
		[]string{"session", "find", "Bash"},
		[]string{"session", "search", "kilo"},
		[]string{"session", "sweep"},
		[]string{"session", "backfill"},
		[]string{"session", "delete", "alfa"},
		[]string{"session", "export", "--all"},
		[]string{"session", "bash-dump"},
		[]string{"message", "read", "1"},
	) {
		c := newMock(t)
		c.Failure = errors.New("foxtrot")
		_, failure, code := run(t, c, arguments...)
		assert.Integer(t, 1, code)
		assert.StringContains(t, "foxtrot", failure)
	}
}

func TestEditPrintsNothingOnSuccess(t *testing.T) {
	c := newMock(t)
	var output, failure string
	var code int
	stray := assert.Capture(
		t,
		func() {
			output, failure, code = run(
				t,
				c,
				"session",
				soil.Edit,
				"alfa",
				"--name",
				"bravo",
			)
		},
	)
	assert.Integer(t, 0, code)
	assert.String(t, "", output)
	assert.String(t, "", failure)
	assert.String(t, "", stray)
}

func TestEditWithoutAChangeIsRefused(t *testing.T) {
	_, failure, code := run(t, newMock(t), "session", soil.Edit, "alfa")
	assert.Integer(t, 1, code)
	assert.StringContains(t, "provide --name or --description", failure)
}

func TestCostEmitsItsWindow(t *testing.T) {
	output, _, code := run(t, newMock(t), "cost", "--days", "7")
	assert.Integer(t, 0, code)
	var r client.CostResponse
	notation.MustDecode(output, &r, false)
	assert.Integer(t, 7, r.Days)
}

func TestPrintIsADocument(t *testing.T) {
	c := newMock(t)
	c.Messages = []client.SessionMessage{
		{Role: "user", Text: "kilo"},
		{Role: "assistant", Text: "lima"},
	}
	output, failure, code := run(t, c, "session", "print", "alfa")
	assert.Integer(t, 0, code)
	assert.String(t, "", failure)
	assert.String(t, "## user [1]\n\nkilo\n\n## assistant\n\nlima\n\n", output)
}

func TestNoUsageYetIsNotFound(t *testing.T) {
	c := newMock(t)
	c.UsageMissing = true
	output, failure, code := run(t, c, "usage")
	assert.Integer(t, 1, code)
	assert.String(t, "", output)
	assert.StringContains(t, "no usage data yet", failure)
}

func TestContextCutsALongMessageAtThreeHundredCharacters(t *testing.T) {
	c := newMock(t)
	long := fmt.Sprintf("x%s", strings.Repeat("ä", 300))
	c.ToolContext = []client.ToolContextResult{
		{
			ToolName: "Bash",
			Before:   &[]client.SessionMessage{{Role: "user", Text: long}},
		},
	}
	output, _, code := run(t, c, "session", "context", "alfa", "Bash")
	assert.Integer(t, 0, code)
	assert.StringContains(
		t,
		fmt.Sprintf("**user:** x%s…\n", strings.Repeat("ä", 299)),
		output,
	)
	assert.True(t, utf8.ValidString(output))
}

func TestAnInvalidTurnRangeIsRefused(t *testing.T) {
	_, failure, code := run(
		t,
		newMock(t),
		"session",
		"print",
		"alfa",
		"--turns",
		"zulu",
	)
	assert.Integer(t, 1, code)
	assert.StringContains(t, "invalid --turns range: zulu", failure)
}
