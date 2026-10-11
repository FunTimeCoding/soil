package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	generative "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/base"
	"testing"
)

func TestSearchConversationsTool(t *testing.T) {
	s := base.New(t)
	s.WriteSearchTranscript("bravo")
	s.Service.CatchUpSearch()
	a := s.NewSession(t)
	a.Announce(a.Name(), "searching")
	result := a.MustCallTool(
		constant.SearchConversations,
		map[string]any{generative.ParameterQuery: "legacy pin"},
	)
	assert.StringContains(t, "(bravo) - 2 hits", result)
	assert.StringContains(
		t,
		"[2026-10-08T01:00:00Z user message u1] where is the legacy protocol pin",
		result,
	)
}

func TestSearchConversationsToolKinds(t *testing.T) {
	s := base.New(t)
	s.WriteSearchTranscript("bravo")
	s.Service.CatchUpSearch()
	a := s.NewSession(t)
	a.Announce(a.Name(), "searching")
	result := a.MustCallTool(
		constant.SearchConversations,
		map[string]any{
			generative.ParameterQuery: "go test",
			constant.Kinds:            []string{"call"},
		},
	)
	assert.StringContains(
		t,
		"[2026-10-08T01:02:00Z assistant call a2] Bash go test ./...",
		result,
	)
}

func TestReadConversationTool(t *testing.T) {
	s := base.New(t)
	s.WriteSearchTranscript("bravo")
	s.Service.CatchUpSearch()
	a := s.NewSession(t)
	a.Announce(a.Name(), "reading")
	result := a.MustCallTool(
		constant.ReadConversation,
		map[string]any{
			constant.Session: "bravo",
			constant.Around:  "a1",
			constant.Count:   1,
		},
	)
	assert.StringContains(t, "[2026-10-08T01:00:00Z user message u1]", result)
	assert.StringContains(
		t,
		"[2026-10-08T01:01:00Z assistant message a1] <- hit\nthe pin lives in setup.go",
		result,
	)
	assert.StringContains(t, "[2026-10-08T01:02:00Z assistant call a2]", result)
}

func TestReadConversationToolUnknownBlock(t *testing.T) {
	s := base.New(t)
	s.WriteSearchTranscript("bravo")
	s.Service.CatchUpSearch()
	a := s.NewSession(t)
	a.Announce(a.Name(), "reading")
	message, e := a.CallToolError(
		constant.ReadConversation,
		map[string]any{constant.Session: "bravo", constant.Around: "nope"},
	)
	assert.FatalOnError(t, e)
	assert.StringContains(t, "no block nope in conversation bravo", message)
}
