package unit

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/message"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/session"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/base"
	"github.com/funtimecoding/soil/pkg/web/markup_tester"
	"net/url"
	"testing"
)

func TestConversationsPageSharesFullWindowFrame(t *testing.T) {
	s := base.New(t)
	page := s.Body(t, constant.ConversationsPath)
	markup_tester.New(t).AssertCount(1, page, "body.full-window")
	markup_tester.New(t).AssertCount(1, page, "nav.container-fluid")
	markup_tester.New(t).AssertCount(1, page, "#notifications")
	markup_tester.New(t).AssertCount(1, page, "link[rel=icon]")
	markup_tester.New(t).AssertText(
		"Select a conversation",
		page,
		"#panel .panel-placeholder",
	)
}

func TestConversationsLinkOpensSessionAndMarksIt(t *testing.T) {
	s := base.New(t)
	s.WriteSessionFile("echo", "echo-slug")
	s.Restart()
	s.Peers.Client.AddSession(session.New("echo"))
	s.Peers.Client.AddMessage(
		"echo",
		message.New("user", "can you help me fix the login bug"),
	)
	page := s.Body(t, constant.ConversationsPath)
	markup_tester.New(t).AssertAttribute(
		"/conversations/echo",
		page,
		"#entry-echo a.entry-link",
		"href",
	)
	markup_tester.New(t).AssertAttribute(
		"true",
		page,
		"#entry-echo a.entry-link",
		"hx-push-url",
	)
	markup_tester.New(t).AssertCount(0, page, ".sidebar-entry[aria-current]")
	opened := s.Body(t, "/conversations/echo")
	markup_tester.New(t).AssertCount(1, opened, "body.full-window")
	markup_tester.New(t).AssertAttribute(
		"true",
		opened,
		"#entry-echo",
		"aria-current",
	)
	markup_tester.New(t).AssertText(
		"can you help me fix the login bug",
		opened,
		"#panel .message:first-child .message-text",
	)
}

func TestConversationsUnknownSessionIsNotFound(t *testing.T) {
	s := base.New(t)
	assert.Integer(t, 404, s.Status(t, "/conversations/nosuch"))
	assert.Integer(t, 404, s.Status(t, "/conversations/nosuch/edit"))
	assert.Integer(
		t,
		404,
		s.PostStatus(
			t,
			"/conversations/nosuch/edit",
			url.Values{constant.Alias: {"renamed"}},
		),
	)
}

func TestConversationsSidebarPagesWhenShown(t *testing.T) {
	s := base.New(t)

	for i := range 31 {
		s.WriteSessionFile(fmt.Sprintf("s%02d", i), fmt.Sprintf("slug-%02d", i))
	}

	s.Restart()
	first := s.Body(t, "/conversations/sidebar")
	markup_tester.New(t).AssertCount(30, first, ".sidebar-entry")
	markup_tester.New(t).AssertAttribute(
		"/conversations/sidebar?offset=30",
		first,
		"small.sidebar-more",
		"hx-get",
	)
	markup_tester.New(t).AssertAttribute(
		"intersect once",
		first,
		"small.sidebar-more",
		"hx-trigger",
	)
	rest := s.Body(t, "/conversations/sidebar?offset=30")
	markup_tester.New(t).AssertCount(1, rest, ".sidebar-entry")
	markup_tester.New(t).AssertCount(0, rest, ".sidebar-more")
}
