package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/chat/teams/user"
	"github.com/funtimecoding/soil/pkg/errors/ambiguous"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gographd/unit/service_tester"
	"testing"
)

func TestUserResolvedByExactNameBeforePrefix(t *testing.T) {
	o := service_tester.New(t)
	o.SeedPeople()
	u, e := o.Service().UserProfile("juliett")
	assert.FatalOnError(t, e)
	assert.String(t, "user-9", u.Identifier)
}

func TestUserResolvedByMail(t *testing.T) {
	o := service_tester.New(t)
	o.SeedPeople()
	u, e := o.Service().UserProfile("Juliett.Romeo@example.com")
	assert.FatalOnError(t, e)
	assert.String(t, "user-8", u.Identifier)
}

func TestUserPrefixMatchingSeveralAmbiguous(t *testing.T) {
	o := service_tester.New(t)
	o.SeedPeople()
	_, e := o.Service().UserProfile("Jul")
	assert.True(t, ambiguous.Is(e))
}

func TestUnknownUserNotFound(t *testing.T) {
	o := service_tester.New(t)
	o.SeedPeople()
	_, e := o.Service().UserProfile("Mike")
	assert.True(t, not_found.Is(e))
}

func TestDirectMessageOpensChatOnce(t *testing.T) {
	o := service_tester.New(t)
	o.SeedPeople()
	_, e := o.Service().SendDirectMessage("kilo@example.com", "hi", nil)
	assert.FatalOnError(t, e)
	sent, f := o.Service().SendDirectMessage("Kilo", "again", nil)
	assert.FatalOnError(t, f)
	assert.String(t, "again", sent.Text)
	chats, g := o.Service().Chats(0)
	assert.FatalOnError(t, g)
	assert.Count(t, 1, chats)
	assert.String(t, "oneOnOne", chats[0].Kind)
	messages, h := o.Service().ChatHistory(chats[0].Identifier, 0)
	assert.FatalOnError(t, h)
	assert.Count(t, 2, messages)
}

func TestDirectMessageTagsMemberOfNewChat(t *testing.T) {
	o := service_tester.New(t)
	o.SeedPeople()
	_, e := o.Service().SendDirectMessage(
		"Kilo",
		"@kilo busy?",
		[]string{"kilo"},
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 1, o.Source.Mentioned)
	assert.String(t, "user-7", o.Source.Mentioned[0].User)
}

func TestGroupChatWithTopic(t *testing.T) {
	o := service_tester.New(t)
	o.SeedPeople()
	h, e := o.Service().CreateChat([]string{"Juliett", "Kilo"}, "LVS")
	assert.FatalOnError(t, e)
	assert.String(t, "group", h.Kind)
	assert.String(t, "LVS", h.Topic)
}

func TestDirectChatWithTopicInvalid(t *testing.T) {
	o := service_tester.New(t)
	o.SeedPeople()
	_, e := o.Service().CreateChat([]string{"Kilo"}, "LVS")
	assert.True(t, validation.Is(e))
}

func TestChatWithOnlyYourselfInvalid(t *testing.T) {
	o := service_tester.New(t)
	o.SeedPeople()
	o.Source.AddUser(user.New("me-1", "Alfa", "alfa@example.com"))
	_, e := o.Service().CreateChat([]string{"alfa@example.com"}, "")
	assert.True(t, validation.Is(e))
}
