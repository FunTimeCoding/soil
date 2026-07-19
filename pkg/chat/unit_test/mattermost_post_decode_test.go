package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/chat/mattermost/constant"
	"github.com/funtimecoding/soil/pkg/chat/mattermost/post"
	"github.com/mattermost/mattermost/server/public/model"
	"testing"
)

func TestDecode(t *testing.T) {
	e := &model.WebSocketEvent{}
	e = e.SetData(map[string]any{constant.PostField: `{"id":"alfa"}`})
	assert.Any(t, &model.Post{Id: "alfa"}, post.Decode(e))
}
