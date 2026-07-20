package coordination

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration_test/base"
	"testing"
)

func TestSummarizePushesIndexer(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "building things")
	a.MustCallTool(
		constant.EditSession,
		map[string]any{
			constant.Alias: "test-session",
		},
	)
	a.MustCallTool(
		constant.Summarize,
		map[string]any{
			constant.Body: "summary for indexing",
		},
	)
	assert.Count(t, 1, s.SummaryIndexer.Pushed)
	assert.String(t, "test-session", s.SummaryIndexer.Pushed[0].Name)
	assert.String(t, "summary for indexing", s.SummaryIndexer.Pushed[0].Body)
}

func TestSummarizeAmendPushesUpdatedBody(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "building things")
	a.MustCallTool(
		constant.EditSession,
		map[string]any{
			constant.Alias: "test-session",
		},
	)
	a.MustCallTool(
		constant.Summarize,
		map[string]any{
			constant.Body: "first draft",
		},
	)
	a.MustCallTool(
		constant.Summarize,
		map[string]any{
			constant.Body: "revised with more detail",
		},
	)
	assert.Count(t, 2, s.SummaryIndexer.Pushed)
	assert.String(
		t,
		"revised with more detail",
		s.SummaryIndexer.Pushed[1].Body,
	)
}
