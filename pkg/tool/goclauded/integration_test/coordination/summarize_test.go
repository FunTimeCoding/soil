package coordination

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/event_query"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration_test/base"
	"testing"
)

func TestSummarize(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "building things")
	a.MustCallTool(
		constant.Summarize,
		map[string]any{
			constant.Body: "built the search index pipeline",
		},
	)
	v := s.Store.ListSummaries()
	assert.Count(t, 1, v)
	assert.String(t, a.Name(), v[0].Name)
	assert.String(t, "built the search index pipeline", v[0].Body)
}

func TestSummarizeHistoryEvent(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "building things")
	a.MustCallTool(
		constant.Summarize,
		map[string]any{
			constant.Body: "session summary content",
		},
	)
	r := a.MustCallTool(constant.History, map[string]any{})
	assert.StringContains(t, "summarized", r)
	assert.StringContains(t, "session summary content", r)
}

func TestSummarizeAmend(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "building things")
	a.MustCallTool(
		constant.Summarize,
		map[string]any{
			constant.Body: "first draft",
		},
	)
	a.MustCallTool(
		constant.Summarize,
		map[string]any{
			constant.Body: "revised summary with more detail",
		},
	)
	v := s.Store.ListSummaries()
	assert.Count(t, 1, v)
	assert.String(t, "revised summary with more detail", v[0].Body)
}

func TestSummarizeAmendOneEvent(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "building things")
	a.MustCallTool(
		constant.Summarize,
		map[string]any{
			constant.Body: "first draft",
		},
	)
	a.MustCallTool(
		constant.Summarize,
		map[string]any{
			constant.Body: "revised",
		},
	)
	v := s.Store.Events(event_query.New().Kind(constant.Summarize).SetLimit(10))
	assert.Count(t, 1, v)
	assert.String(t, "revised", v[0].Metadata[constant.Body])
}

func TestSummarizeBeforeAnnounce(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	assert.StringContains(
		t,
		"announce first",
		a.MustCallToolError(
			constant.Summarize,
			map[string]any{
				constant.Body: "should fail",
			},
		),
	)
}
