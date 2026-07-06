package post

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/chat/mattermost/constant"
	"github.com/mattermost/mattermost/server/public/model"
	"testing"
)

func TestDecode(t *testing.T) {
	e := &model.WebSocketEvent{}
	e = e.SetData(map[string]any{constant.PostField: `{"id":"alfa"}`})
	assert.Any(t, &model.Post{Id: "alfa"}, Decode(e))
}
