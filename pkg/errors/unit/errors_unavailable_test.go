package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"testing"
	"time"
)

func TestAnOutageKeepsItsStartAndReason(t *testing.T) {
	since := time.Date(2026, 1, 2, 3, 4, 5, 0, time.UTC)
	e := unavailable.New("alfa", since, "not connected")
	assert.String(
		t,
		"alfa unreachable since 2026-01-02 03:04:05: not connected",
		e.Error(),
	)
	assert.True(t, since.Equal(e.Since))
	assert.String(t, "not connected", e.Reason)
}

func TestAFormattedOutageHasNoStart(t *testing.T) {
	e := unavailable.Format("bravo is down")
	assert.String(t, "bravo is down", e.Error())
	assert.True(t, e.Since.IsZero())
	assert.String(t, "", e.Reason)
}
