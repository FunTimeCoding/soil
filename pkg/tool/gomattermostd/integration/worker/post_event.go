package worker

import (
	"github.com/funtimecoding/soil/pkg/chat/constant"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/mattermost/mattermost/server/public/model"
)

func postEvent(p *model.Post) *model.WebSocketEvent {
	v := &model.WebSocketEvent{}
	v = v.SetEvent(model.WebsocketEventPosted)

	return v.SetData(
		map[string]any{constant.MattermostPostField: notation.Encode(p, false)},
	)
}
