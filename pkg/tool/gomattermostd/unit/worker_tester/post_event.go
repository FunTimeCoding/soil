package worker_tester

import (
	chatConstant "github.com/funtimecoding/soil/pkg/chat/constant"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/mattermost/mattermost/server/public/model"
)

func PostEvent(p *model.Post) *model.WebSocketEvent {
	v := &model.WebSocketEvent{}
	v = v.SetEvent(model.WebsocketEventPosted)

	return v.SetData(
		map[string]any{
			chatConstant.MattermostPostField: notation.Encode(p, false),
		},
	)
}
