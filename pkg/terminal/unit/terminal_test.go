package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/console/response"
	"github.com/funtimecoding/soil/pkg/terminal/terminal_tester"
	"net/http"
	"testing"
)

func TestNoticelnWritesToFailureWithoutEndingCommand(t *testing.T) {
	f := terminal_tester.New()
	f.Terminal.Line("alfa")
	f.Terminal.Noticeln("bravo", 2)
	assert.String(t, "alfa\n", f.Output.String())
	assert.String(t, "bravo 2\n", f.Failure.String())
	assert.Count(t, 0, f.Exits)
	assert.Count(t, 0, f.Ends.Outcomes)
}

func TestSuccessfulResponseIsPrintedWithoutExiting(t *testing.T) {
	f := terminal_tester.New()
	f.Terminal.Emit(response.New("alfa", http.StatusOK))
	assert.String(t, "alfa\n", f.Output.String())
	assert.String(t, "", f.Failure.String())
	assert.Count(t, 0, f.Exits)
	assert.Count(t, 0, f.Ends.Outcomes)
}

func TestRejectedResponseGoesToFailureAndExitsOne(t *testing.T) {
	f := terminal_tester.New()
	f.Terminal.Emit(response.New("not found", http.StatusNotFound))
	assert.String(t, "", f.Output.String())
	assert.String(t, "not found\n", f.Failure.String())
	assert.Integers(t, []int{1}, f.Exits)
	assert.Strings(t, []string{"error"}, f.Ends.Outcomes)
}

func TestExitfWritesToFailureAndExitsOne(t *testing.T) {
	f := terminal_tester.New()
	f.Terminal.Exitf("invalid level: %s\n", "alfa")
	assert.String(t, "invalid level: alfa\n", f.Failure.String())
	assert.Integers(t, []int{1}, f.Exits)
	assert.Strings(t, []string{"error"}, f.Ends.Outcomes)
}

func TestExitZeroEndsCommandAsSucceeded(t *testing.T) {
	f := terminal_tester.New()
	f.Terminal.Exit(0)
	assert.Integers(t, []int{0}, f.Exits)
	assert.Strings(t, []string{"success"}, f.Ends.Outcomes)
}

func TestBlocklnWritesToFailureAndEndsCommandAsBlocked(t *testing.T) {
	f := terminal_tester.New()
	f.Terminal.Blockln(2, "alfa")
	assert.String(t, "", f.Output.String())
	assert.String(t, "alfa\n", f.Failure.String())
	assert.Integers(t, []int{2}, f.Exits)
	assert.Strings(t, []string{"blocked"}, f.Ends.Outcomes)
}

func TestRejectJoinsStatusAndBody(t *testing.T) {
	f := terminal_tester.New()
	f.Terminal.Reject("401 Unauthorized", []byte("alfa"))
	assert.String(t, "401 Unauthorized: alfa\n", f.Failure.String())
	assert.Integers(t, []int{1}, f.Exits)
	assert.Strings(t, []string{"error"}, f.Ends.Outcomes)
}

func TestRejectWithoutBodyPrintsStatus(t *testing.T) {
	f := terminal_tester.New()
	f.Terminal.Reject("502 Bad Gateway", nil)
	assert.String(t, "502 Bad Gateway\n", f.Failure.String())
	assert.Integers(t, []int{1}, f.Exits)
}

func TestRequiredReturnsSetVariable(t *testing.T) {
	t.Setenv("TERMINAL_ALFA", "bravo")
	f := terminal_tester.New()
	assert.String(t, "bravo", f.Terminal.Required("TERMINAL_ALFA"))
	assert.Count(t, 0, f.Exits)
}

func TestRequiredExitsOneWhenVariableIsUnset(t *testing.T) {
	t.Setenv("TERMINAL_ALFA", "")
	f := terminal_tester.New()
	f.Terminal.Required("TERMINAL_ALFA")
	assert.String(t, "", f.Output.String())
	assert.String(t, "TERMINAL_ALFA not set\n", f.Failure.String())
	assert.Integers(t, []int{1}, f.Exits)
	assert.Strings(t, []string{"error"}, f.Ends.Outcomes)
}

func TestEmptyBodyPrintsNothing(t *testing.T) {
	f := terminal_tester.New()
	f.Terminal.Emit(response.New("", http.StatusNoContent))
	assert.String(t, "", f.Output.String())
}
