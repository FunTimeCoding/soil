package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	chatConstant "github.com/funtimecoding/soil/pkg/chat/constant"
	"github.com/funtimecoding/soil/pkg/chat/mattermost/post"
	"github.com/funtimecoding/soil/pkg/chat/mattermost/reaction"
	"github.com/funtimecoding/soil/pkg/chat/telegram/message"
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/mattermost/mattermost/server/public/model"
	"testing"
)

func TestClient(t *testing.T) {
	assert.String(t, "MATTERMOST_HOST", chatConstant.MattermostHostEnvironment)
	assert.String(
		t,
		"MATTERMOST_TOKEN",
		chatConstant.MattermostTokenEnvironment,
	)
	assert.String(t, "MATTERMOST_TEAM", chatConstant.MattermostTeamEnvironment)
	assert.String(
		t,
		"MATTERMOST_CHANNEL",
		chatConstant.MattermostChannelEnvironment,
	)
	assert.String(
		t,
		"MATTERMOST_INSECURE",
		chatConstant.MattermostInsecureEnvironment,
	)
	assert.String(t, "construction", chatConstant.MattermostConstruction)
	assert.String(t, "hourglass_flowing_sand", chatConstant.MattermostHourglass)
	assert.String(t, "repeat", chatConstant.MattermostRepeat)
	assert.String(t, "thread", chatConstant.MattermostThread)
}

func TestConstant(t *testing.T) {
	assert.String(t, "TELEGRAM_TOKEN", chatConstant.TelegramTokenEnvironment)
}

func TestDecode(t *testing.T) {
	e := &model.WebSocketEvent{}
	e = e.SetData(
		map[string]any{chatConstant.MattermostPostField: `{"id":"alfa"}`},
	)
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
			chatConstant.MattermostReactionField: `{"post_id":"alfa","emoji_name":"bravo"}`,
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
				From: &tgbotapi.User{UserName: constant.UpperAlfa},
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
