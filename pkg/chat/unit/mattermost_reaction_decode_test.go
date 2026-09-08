package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/chat/constant"
	"github.com/funtimecoding/soil/pkg/chat/mattermost/reaction"
	"github.com/mattermost/mattermost/server/public/model"
	"testing"
)

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
