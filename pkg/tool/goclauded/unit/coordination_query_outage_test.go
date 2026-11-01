package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/service_tester"
	"testing"
)

func TestACompletionAnswersWhileGoquerydIsAway(t *testing.T) {
	s := base.New(t)
	a := s.NewSession(t)
	a.Announce(a.Name(), "building things")
	a.MustCallTool(
		constant.EditSession,
		map[string]any{constant.Alias: "test-completion"},
	)
	s.Query.Failure = unavailable.Format("goqueryd down")
	a.MustCallTool(
		constant.Complete,
		map[string]any{constant.Message: "built the thing"},
	)
	assert.Count(t, 0, s.Pushed(constant.CompletionCollection))
	assert.Count(t, 0, s.Reporter.Events())
}

func TestTheQueryGuardHoldsCallsOffAfterAnOutage(t *testing.T) {
	s := base.New(t)
	a := s.NewSession(t)
	a.Announce(a.Name(), "building things")
	a.MustCallTool(
		constant.EditSession,
		map[string]any{constant.Alias: "test-completion"},
	)
	s.Query.Failure = unavailable.Format("goqueryd down")
	a.MustCallTool(
		constant.Complete,
		map[string]any{constant.Message: "built the thing"},
	)
	s.Query.Failure = nil
	a.MustCallTool(
		constant.Summarize,
		map[string]any{constant.Body: "the summary"},
	)
	assert.Count(t, 0, s.Pushed(constant.SummaryCollection))
}

func TestReconcileCatchesUpACompletionAWriteMissed(t *testing.T) {
	s := base.New(t)
	a := s.NewSession(t)
	a.Announce(a.Name(), "building things")
	a.MustCallTool(
		constant.EditSession,
		map[string]any{constant.Alias: "test-completion"},
	)
	s.Query.Failure = validation.New("refused")
	a.MustCallTool(
		constant.Complete,
		map[string]any{constant.Message: "built the thing"},
	)
	assert.Count(t, 0, s.Pushed(constant.CompletionCollection))
	s.Query.Failure = nil
	assert.Count(t, 0, s.Service.Reconcile(t.Context()))
	pushed := s.Pushed(constant.CompletionCollection)
	assert.Count(t, 1, pushed)
	assert.String(t, "test-completion/1", pushed[0].Path)
}

func TestCatchUpSearchStopsWhenCancelled(t *testing.T) {
	s := service_tester.New(t)
	s.WriteSessionFile("alfa", "alfa-slug")
	x, cancel := context.WithCancel(t.Context())
	cancel()
	assert.Count(t, 0, s.Service.CatchUpSearch(x))
	r, e := s.Search.Search("login bug", nil, 20)
	assert.FatalOnError(t, e)
	assert.Count(t, 0, r)
}
