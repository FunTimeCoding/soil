package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/types/argument"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/service_tester"
	"testing"
)

func TestRedactionFailsClosedWhileMemoryIsDown(t *testing.T) {
	s := service_tester.New(t)
	s.Store.EnsureSession("closed-session")
	s.WriteContextLoadFile("closed-session")
	s.Transcript.EnrichSession("closed-session")
	s.MemoryPeer.Failure = refusedConnection()
	loads, e := s.Transcript.ContextLoadsBySession(
		t.Context(),
		"closed-session",
	)
	assert.Count(t, 0, loads)
	assert.True(t, unavailable.Is(e))
	assert.StringContains(t, "memory unreachable since", e.Error())
	assert.StringContains(t, "loads hidden", e.Error())
}

func TestTheContextPanelSaysWhenMemoryIsDown(t *testing.T) {
	s := service_tester.NewPanel(t, "panel-session")
	s.MemoryPeer.Failure = refusedConnection()
	assert.StringContains(
		t,
		"memory unreachable since",
		s.Memory.ContextPanel(t.Context(), "panel-session"),
	)
}

func TestTheTimelineSaysWhenMemoryIsDown(t *testing.T) {
	s := service_tester.New(t)
	s.MemoryPeer.Failure = refusedConnection()
	entries, e := s.Event.Timeline(t.Context(), argument.Timeline{Limit: 10})
	assert.FatalOnError(t, e)
	assert.Count(t, 1, entries)
	assert.String(t, "memory_unreachable", entries[0].Kind)
}

func TestAMomentWhileMemoryIsDownIsDeliveredLaterWithItsTime(t *testing.T) {
	s := service_tester.New(t)
	s.Store.EnsureSession("moment-session")
	s.MemoryPeer.Failure = refusedConnection()
	s.Moment("moment-session", "kilo", "the light through the blinds")
	assert.Count(t, 0, s.MemoryPeer.Impressions)
	pending, e := s.Store.Store.PendingImpressions()
	assert.FatalOnError(t, e)
	assert.Count(t, 1, pending)
	s.MemoryPeer.Failure = nil
	assert.FatalOnError(t, s.MemoryGuard().Wait(t.Context()))
	assert.Count(t, 0, s.Memory.DeliverImpressions())
	assert.Count(t, 1, s.MemoryPeer.Impressions)
	assert.String(
		t,
		"the light through the blinds",
		s.MemoryPeer.Impressions[0].Content,
	)
	assert.True(
		t,
		s.MemoryPeer.Impressions[0].CapturedAt.Equal(pending[0].CapturedAt),
	)
	left, f := s.Store.Store.PendingImpressions()
	assert.FatalOnError(t, f)
	assert.Count(t, 0, left)
	assert.Count(t, 0, s.Reporter.Events())
}

func TestARefusedImpressionGoesToSentryAndLeavesTheOutbox(t *testing.T) {
	s := service_tester.New(t)
	s.Store.EnsureSession("refused-session")
	s.MemoryPeer.Failure = refusedConnection()
	s.Moment("refused-session", "lima", "first")
	s.Moment("refused-session", "lima", "second")
	s.MemoryPeer.Failure = nil
	assert.FatalOnError(t, s.MemoryGuard().Wait(t.Context()))
	s.MemoryPeer.Failure = errors.New("save impression status: 400: too long")
	assert.Count(t, 0, s.Memory.DeliverImpressions())
	left, e := s.Store.Store.PendingImpressions()
	assert.FatalOnError(t, e)
	assert.Count(t, 0, left)
	reported := s.Reporter.Events()
	assert.Count(t, 2, reported)
	assert.MapValue(
		t,
		"second",
		reported[1].Context,
		constant.ImpressionContent,
	)
}
