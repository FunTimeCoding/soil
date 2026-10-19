package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/classify"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/errors/unreachable"
	"testing"
)

func TestCaptureShowsTheContextAndTheFault(t *testing.T) {
	m := memory.New()
	text, _ := classify.Capture(m, errors.New("disk full"), "load profile")
	assert.String(t, "load profile: disk full", text)
	assert.Count(t, 1, m.Events())
	assert.String(t, "disk full", m.Events()[0].Error.Error())
}

func TestCaptureNamesAKnownFault(t *testing.T) {
	m := memory.New()
	text, _ := classify.Capture(m, unreachable.Format("store unreachable"), "")
	assert.String(t, "store unreachable", text)
	assert.String(t, "store unreachable", m.Events()[0].Error.Error())
}

func TestCaptureTitlesAQuietClassAsAMissingBranch(t *testing.T) {
	m := memory.New()
	text, _ := classify.Capture(
		m,
		not_found.Format("memory not found: 12"),
		"load profile",
	)
	assert.String(t, "load profile: memory not found: 12", text)
	assert.Count(t, 1, m.Events())
	assert.String(
		t,
		"missing not_found branch: memory not found: 12",
		m.Events()[0].Error.Error(),
	)
}
