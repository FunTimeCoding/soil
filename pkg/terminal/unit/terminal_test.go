package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/console/response"
	"net/http"
	"testing"
)

func TestASuccessfulResponseIsPrintedWithoutExiting(t *testing.T) {
	f := newFixture()
	f.terminal.Emit(response.New("alfa", http.StatusOK))
	assert.String(t, "alfa\n", f.output.String())
	assert.String(t, "", f.failure.String())
	assert.Count(t, 0, f.exits)
	assert.Count(t, 0, f.ends.outcomes)
}

func TestARejectedResponseGoesToFailureAndExitsOne(t *testing.T) {
	f := newFixture()
	f.terminal.Emit(response.New("not found", http.StatusNotFound))
	assert.String(t, "", f.output.String())
	assert.String(t, "not found\n", f.failure.String())
	assert.Integers(t, []int{1}, f.exits)
	assert.Strings(t, []string{"error"}, f.ends.outcomes)
}

func TestExitfWritesToFailureAndExitsOne(t *testing.T) {
	f := newFixture()
	f.terminal.Exitf("invalid level: %s\n", "alfa")
	assert.String(t, "invalid level: alfa\n", f.failure.String())
	assert.Integers(t, []int{1}, f.exits)
	assert.Strings(t, []string{"error"}, f.ends.outcomes)
}

func TestExitZeroEndsTheCommandAsSucceeded(t *testing.T) {
	f := newFixture()
	f.terminal.Exit(0)
	assert.Integers(t, []int{0}, f.exits)
	assert.Strings(t, []string{"success"}, f.ends.outcomes)
}

func TestBlocklnWritesToFailureAndEndsTheCommandAsBlocked(t *testing.T) {
	f := newFixture()
	f.terminal.Blockln(2, "alfa")
	assert.String(t, "", f.output.String())
	assert.String(t, "alfa\n", f.failure.String())
	assert.Integers(t, []int{2}, f.exits)
	assert.Strings(t, []string{"blocked"}, f.ends.outcomes)
}

func TestRejectJoinsStatusAndBody(t *testing.T) {
	f := newFixture()
	f.terminal.Reject("401 Unauthorized", []byte("alfa"))
	assert.String(t, "401 Unauthorized: alfa\n", f.failure.String())
	assert.Integers(t, []int{1}, f.exits)
	assert.Strings(t, []string{"error"}, f.ends.outcomes)
}

func TestRejectWithoutBodyPrintsTheStatus(t *testing.T) {
	f := newFixture()
	f.terminal.Reject("502 Bad Gateway", nil)
	assert.String(t, "502 Bad Gateway\n", f.failure.String())
	assert.Integers(t, []int{1}, f.exits)
}

func TestRequiredReturnsASetVariable(t *testing.T) {
	t.Setenv("TERMINAL_ALFA", "bravo")
	f := newFixture()
	assert.String(t, "bravo", f.terminal.Required("TERMINAL_ALFA"))
	assert.Count(t, 0, f.exits)
}

func TestRequiredExitsOneWhenTheVariableIsUnset(t *testing.T) {
	t.Setenv("TERMINAL_ALFA", "")
	f := newFixture()
	f.terminal.Required("TERMINAL_ALFA")
	assert.String(t, "", f.output.String())
	assert.String(t, "TERMINAL_ALFA not set\n", f.failure.String())
	assert.Integers(t, []int{1}, f.exits)
	assert.Strings(t, []string{"error"}, f.ends.outcomes)
}

func TestAnEmptyBodyPrintsNothing(t *testing.T) {
	f := newFixture()
	f.terminal.Emit(response.New("", http.StatusNoContent))
	assert.String(t, "", f.output.String())
}
