package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/chat/constant"
	"github.com/funtimecoding/soil/pkg/chat/mattermost/post"
	"github.com/funtimecoding/soil/pkg/chat/mattermost/reaction"
	"github.com/funtimecoding/soil/pkg/chat/telegram/message"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/mattermost/mattermost/server/public/model"
	"testing"
)

func TestClient(t *testing.T) {
	assert.String(t, "MATTERMOST_HOST", constant.MattermostHostEnvironment)
	assert.String(t, "MATTERMOST_TOKEN", constant.MattermostTokenEnvironment)
	assert.String(t, "MATTERMOST_TEAM", constant.MattermostTeamEnvironment)
	assert.String(
		t,
		"MATTERMOST_CHANNEL",
		constant.MattermostChannelEnvironment,
	)
	assert.String(
		t,
		"MATTERMOST_INSECURE",
		constant.MattermostInsecureEnvironment,
	)
	assert.String(t, "construction", constant.MattermostProgressReaction)
	assert.String(
		t,
		"hourglass_flowing_sand",
		constant.MattermostWaitingReaction,
	)
	assert.String(t, "repeat", constant.MattermostForwardedReaction)
	assert.String(t, "thread", constant.MattermostThreadReaction)
}

func TestConstant(t *testing.T) {
	assert.String(t, "TELEGRAM_TOKEN", constant.TelegramTokenEnvironment)
}

func TestDecode(t *testing.T) {
	e := &model.WebSocketEvent{}
	e = e.SetData(map[string]any{constant.MattermostPostField: `{"id":"alfa"}`})
	assert.Any(t, &model.Post{Id: "alfa"}, post.Decode(e))
}

func TestFromList(t *testing.T) {
	bravo := &model.Post{Id: "bravo", CreateAt: 3000}
	alfa := &model.Post{Id: "alfa", CreateAt: 2000}
	parent := &model.Post{Id: "parent", CreateAt: 1000}
	l := &model.PostList{
		Order: []string{"bravo", "alfa", "missing"},
		Posts: map[string]*model.Post{
			"alfa":   alfa,
			"bravo":  bravo,
			"parent": parent,
		},
	}
	assert.Any(t, []*model.Post{bravo, alfa}, post.FromList(l, false))
	assert.Any(t, []*model.Post{alfa, bravo}, post.FromList(l, true))
}

func TestPost(t *testing.T) {
	assert.NotNil(t, post.New(&model.Post{}))
}

func TestMattermostReactionDecode(t *testing.T) {
	e := &model.WebSocketEvent{}
	e = e.SetData(
		map[string]any{
			constant.MattermostReactionField: `{"post_id":"alfa","emoji_name":"bravo"}`,
		},
	)
	assert.Any(
		t,
		&model.Reaction{PostId: "alfa", EmojiName: "bravo"},
		reaction.Decode(e),
	)
}

func TestMessage(t *testing.T) {
	assert.NotNil(
		t,
		message.New(
			&tgbotapi.Message{
				From: &tgbotapi.User{UserName: stringsConstant.UpperAlfa},
			},
		),
	)
}

func TestNewSlice(t *testing.T) {
	assert.Count(t, 0, message.NewSlice([]*tgbotapi.Message{}))
}

func TestStoreRoundTrip(t *testing.T) {
	s := newStore(t)
	s.MustSaveChannel(-100, "announcements")
	s.MustSaveChannel(-100, "announcements")
	s.MustSaveUser(7, "admin")
	assert.Integer(t, 1, len(s.MustChannels()))
	assert.Integer(t, 1, len(s.MustUsers()))
	h := s.MustChannelByName("announcements")
	assert.NotNil(t, h)
	assert.Integer(t, -100, h.Identifier)
	assert.True(t, s.MustChannelByName("missing") == nil)
}
