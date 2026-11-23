package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/chat/teams/chat"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gographd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gographd/unit/service_tester"
	"testing"
)

func TestChannelHistoryByNames(t *testing.T) {
	o := service_tester.New(t)
	o.Seed()
	messages, e := o.Service().ChannelHistory("Security", "cve-2026-1", 0)
	assert.FatalOnError(t, e)
	assert.Count(t, 1, messages)
	assert.String(t, "patched?", messages[0].Text)
}

func TestUnknownChannelNotFound(t *testing.T) {
	o := service_tester.New(t)
	o.Seed()
	_, e := o.Service().ChannelHistory("Security", "Random", 0)
	assert.True(t, not_found.Is(e))
	assert.String(t, "channel not found: Random", e.Error())
}

func TestThreadRepliesByIdentifier(t *testing.T) {
	o := service_tester.New(t)
	o.Seed()
	replies, e := o.Service().ThreadReplies("team-1", "channel-2", "100", 0)
	assert.FatalOnError(t, e)
	assert.Count(t, 1, replies)
	assert.String(t, "yes", replies[0].Text)
}

func TestNonNumericMessageInvalid(t *testing.T) {
	o := service_tester.New(t)
	o.Seed()
	_, e := o.Service().ThreadReplies("team-1", "channel-2", "../100", 0)
	assert.True(t, validation.Is(e))
}

func TestChatHistoryByIdentifier(t *testing.T) {
	o := service_tester.New(t)
	o.Seed()
	messages, e := o.Service().ChatHistory("19:alpha@thread.v2", 0)
	assert.FatalOnError(t, e)
	assert.Count(t, 1, messages)
	assert.String(t, "Knut", messages[0].Author)
}

func TestChatWithPathStepInvalid(t *testing.T) {
	o := service_tester.New(t)
	o.Seed()
	_, e := o.Service().ChatHistory("19:alpha@thread.v2/../x", 0)
	assert.True(t, validation.Is(e))
	_, f := o.Service().ChatHistory(constant.ParentStep, 0)
	assert.True(t, validation.Is(f))
	_, g := o.Service().ChatHistory("19:alpha@thread.v2?top=1", 0)
	assert.True(t, validation.Is(g))
}

func TestChatsLimited(t *testing.T) {
	o := service_tester.New(t)
	o.Seed()
	o.Source.AddChat(chat.New("19:beta@thread.v2", "oneOnOne"))
	chats, e := o.Service().Chats(1)
	assert.FatalOnError(t, e)
	assert.Count(t, 1, chats)
}
