package unit

import (
	"errors"
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/caller_left"
	"github.com/funtimecoding/soil/pkg/errors/classify"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"testing"
)

func TestACallerThatLeftIsNotCaptured(t *testing.T) {
	r := memory.New()
	text, event := classify.Capture(
		t.Context(),
		r,
		caller_left.New(errors.New("context canceled")),
		"",
	)
	assert.StringContains(t, "context canceled", text)
	assert.String(t, "", event)
	assert.Count(t, 0, r.Events())
}

func TestAWrappedCallerThatLeftIsNotCaptured(t *testing.T) {
	r := memory.New()
	_, event := classify.Capture(
		t.Context(),
		r,
		fmt.Errorf("probe: %w", caller_left.New(errors.New("context canceled"))),
		"",
	)
	assert.String(t, "", event)
	assert.Count(t, 0, r.Events())
}

func TestAFaultIsStillCaptured(t *testing.T) {
	r := memory.New()
	_, event := classify.Capture(t.Context(), r, errors.New("disk full"), "")
	assert.String(t, "memory-1", event)
	assert.Count(t, 1, r.Events())
}
