package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/base"
	"github.com/funtimecoding/soil/pkg/web/unit/markup_tester"
	"strings"
	"testing"
)

func TestConversationsSearchResults(t *testing.T) {
	s := base.New(t)
	s.WriteSearchTranscript("delta")
	s.Tick()
	page := body(t, s.Server, "/conversations/search?query=legacy+pin")
	assert.StringContains(t, `class="search-result"`, page)
	assert.StringContains(t, "2 hits", page)
	assert.StringContains(t, `<mark class="highlight">legacy</mark>`, page)
	assert.StringContains(t, "/conversations/delta/hits?around=u1", page)
}

func TestConversationsSearchAddressRestoresSearch(t *testing.T) {
	s := base.New(t)
	s.WriteSearchTranscript("delta")
	s.Tick()
	page := body(
		t,
		s.Server,
		"/conversations/search?query=legacy+pin&kind=message&kind=call",
	)
	markup_tester.AssertCount(t, 1, page, "body.full-window")
	markup_tester.AssertAttribute(
		t,
		"legacy pin",
		page,
		".search-form input[name=query]",
		constant.Value,
	)
	markup_tester.AssertCount(
		t,
		2,
		page,
		".search-form input[name=kind][checked]",
	)
	markup_tester.AssertCount(
		t,
		0,
		page,
		`.search-form input[name=kind][value=edit][checked]`,
	)
	markup_tester.AssertCount(t, 1, page, "#sidebar-entries .search-result")
	markup_tester.AssertAttribute(
		t,
		"true",
		page,
		".search-form",
		"hx-push-url",
	)
}

func TestConversationsSearchAnswersExtendedRequestWithResultsOnly(t *testing.T) {
	s := base.New(t)
	s.WriteSearchTranscript("delta")
	s.Tick()
	answer := extendedBody(
		t,
		s.Server,
		"/conversations/search?query=legacy+pin",
	)
	assert.StringContains(t, `class="search-result"`, answer)
	assert.True(t, !strings.Contains(answer, "<html"))
}

func TestConversationsHitAddressRestoresSearchAndHits(t *testing.T) {
	s := base.New(t)
	s.WriteSearchTranscript("delta")
	s.Tick()
	page := body(
		t,
		s.Server,
		"/conversations/delta/hits?query=pin&kind=message&around=a1",
	)
	markup_tester.AssertCount(t, 1, page, "#sidebar-entries .search-result")
	markup_tester.AssertCount(t, 1, page, "#panel #block-a1.search-current")
	markup_tester.AssertAttribute(
		t,
		"true",
		page,
		"a.search-snippet",
		"hx-push-url",
	)
}

func TestConversationsSearchEmptyQueryListsSessions(t *testing.T) {
	s := base.New(t)
	s.WriteSessionFile("echo", "echo-slug")
	s.Restart()
	page := body(t, s.Server, "/conversations/search?query=")
	markup_tester.AssertCount(t, 1, page, "#sidebar-entries .sidebar-entry")
	markup_tester.AssertCount(t, 0, page, "#sidebar-entries .search-result")
}

func TestConversationsHitsPanel(t *testing.T) {
	s := base.New(t)
	s.WriteSearchTranscript("delta")
	s.Tick()
	page := body(
		t,
		s.Server,
		"/conversations/delta/hits?query=pin+go&kind=message&kind=call&around=a1",
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
	assert.StringContains(t, `<mark class="highlight">pin</mark>`, page)
	assert.StringContains(t, `id="hit-counter"`, page)
}

func TestConversationsPageCarriesSearchWiring(t *testing.T) {
	s := base.New(t)
	page := body(t, s.Server, constant.ConversationsPath)
	assert.StringContains(t, `hx-sync="this:replace"`, page)
	assert.StringContains(t, "input delay:200ms, submit", page)
	assert.StringContains(t, `name="kind"`, page)
	assert.StringContains(t, "search-current", page)
}
