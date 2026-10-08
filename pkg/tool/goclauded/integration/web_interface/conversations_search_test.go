package web_interface

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration/base"
	"strings"
	"testing"
)

func TestConversationsSearchResults(t *testing.T) {
	s := base.New(t)
	s.WriteSearchTranscript("delta")
	s.Service.CatchUpSearch()
	page := body(
		t,
		fmt.Sprintf(
			"http://localhost:%d/conversations/search?query=legacy+pin",
			s.Port,
		),
	)
	assert.StringContains(t, `class="search-result"`, page)
	assert.StringContains(t, "2 hits", page)
	assert.StringContains(t, "<mark>legacy</mark>", page)
	assert.StringContains(t, "/conversations/delta/hits?around=u1", page)
}

func TestConversationsSearchEmptyQueryListsSessions(t *testing.T) {
	s := base.New(t)
	s.WriteSessionFile("echo", "echo-slug")
	s.Service.PopulateCache()
	s.Service.BackfillSessions()
	s.Service.CheckConsistency()
	page := body(
		t,
		fmt.Sprintf("http://localhost:%d/conversations/search?query=", s.Port),
	)
	assert.StringContains(t, "sidebar-entry", page)
	assert.True(t, !strings.Contains(page, "search-result"))
}

func TestConversationsHitsPanel(t *testing.T) {
	s := base.New(t)
	s.WriteSearchTranscript("delta")
	s.Service.CatchUpSearch()
	page := body(
		t,
		fmt.Sprintf(
			"http://localhost:%d/conversations/delta/hits?query=pin+go&kind=message&kind=call&around=a1",
			s.Port,
		),
	)
	assert.StringContains(t, `id="block-a1"`, page)
	assert.StringContains(
		t,
		`class="message message-assistant search-hit search-current"`,
		page,
	)
	assert.StringContains(
		t,
		`class="message message-assistant message-tool search-hit"`,
		page,
	)
	assert.StringContains(t, "<mark>pin</mark>", page)
	assert.StringContains(t, `id="hit-counter"`, page)
}

func TestConversationsPageCarriesSearchWiring(t *testing.T) {
	s := base.New(t)
	page := body(t, fmt.Sprintf("http://localhost:%d/conversations", s.Port))
	assert.StringContains(t, `hx-sync="this:replace"`, page)
	assert.StringContains(t, "input delay:200ms, submit", page)
	assert.StringContains(t, `name="kind"`, page)
	assert.StringContains(t, "search-current", page)
}
