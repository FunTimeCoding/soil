package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/base"
	"testing"
)

func TestSummarizePushesIndexer(t *testing.T) {
	s := base.New(t)
	a := s.NewSession(t)
	a.Announce(a.Name(), "building things")
	a.MustCallTool(
		constant.EditSession,
		map[string]any{constant.Alias: "test-session"},
	)
	a.MustCallTool(
		constant.Summarize,
		map[string]any{constant.Body: "summary for indexing"},
	)
	pushed := s.Pushed(constant.SummaryCollection)
	assert.Count(t, 1, pushed)
	assert.String(t, "test-session", pushed[0].Path)
	assert.String(t, "summary for indexing", pushed[0].Body)
	assert.String(t, "session-summary", *pushed[0].SourceType)
}

func TestSummarizeAmendPushesUpdatedBody(t *testing.T) {
	s := base.New(t)
	a := s.NewSession(t)
	a.Announce(a.Name(), "building things")
	a.MustCallTool(
		constant.EditSession,
		map[string]any{constant.Alias: "test-session"},
	)
	a.MustCallTool(
		constant.Summarize,
		map[string]any{constant.Body: "first draft"},
	)
	a.MustCallTool(
		constant.Summarize,
		map[string]any{constant.Body: "revised with more detail"},
	)
	pushed := s.Pushed(constant.SummaryCollection)
	assert.Count(t, 2, pushed)
	assert.String(t, "revised with more detail", pushed[1].Body)
}
