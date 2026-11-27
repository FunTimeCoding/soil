package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/chat/teams/member"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gographd/unit/service_tester"
	"testing"
)

func TestChannelMessagePostedByNames(t *testing.T) {
	o := service_tester.New(t)
	o.Seed()
	sent, e := o.Service().PostChannelMessage("kitchen", "Lunch", "status", nil)
	assert.FatalOnError(t, e)
	assert.String(t, "status", sent.Text)
	messages, f := o.Service().ChannelHistory("Kitchen", "Lunch", 0)
	assert.FatalOnError(t, f)
	assert.Count(t, 2, messages)
}

func TestThreadReplyPosted(t *testing.T) {
	o := service_tester.New(t)
	o.Seed()
	_, e := o.Service().ReplyToThread("Kitchen", "Lunch", "100", "done", nil)
	assert.FatalOnError(t, e)
	replies, f := o.Service().ThreadReplies("Kitchen", "Lunch", "100", 0)
	assert.FatalOnError(t, f)
	assert.Count(t, 2, replies)
}

func TestThreadReplyToMessageOfOtherShapeInvalid(t *testing.T) {
	o := service_tester.New(t)
	o.Seed()
	_, e := o.Service().ReplyToThread("Kitchen", "Lunch", "../100", "done", nil)
	assert.True(t, validation.Is(e))
}

func TestThreadReplyToUnknownMessageNotFound(t *testing.T) {
	o := service_tester.New(t)
	o.Seed()
	_, e := o.Service().ReplyToThread("Kitchen", "Lunch", "999", "done", nil)
	assert.True(t, not_found.Is(e))
}

func TestChannelMentionResolvedAgainstChannelMembers(t *testing.T) {
	o := service_tester.New(t)
	o.Seed()
	o.Source.AddChannelMember(
		"team-1",
		"channel-2",
		member.New("Kilo", "user-7"),
	)
	o.Source.AddChatMember("19:alpha@thread.v2", member.New("Lima", "user-6"))
	_, e := o.Service().ReplyToThread(
		"Kitchen",
		"Lunch",
		"100",
		"@kilo done",
		[]string{"kilo"},
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 1, o.Source.Mentioned)
	assert.String(t, "user-7", o.Source.Mentioned[0].User)
	_, f := o.Service().ReplyToThread(
		"Kitchen",
		"Lunch",
		"100",
		"@Lima done",
		[]string{"Lima"},
	)
	assert.True(t, not_found.Is(f))
}
