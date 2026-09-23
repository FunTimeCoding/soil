package worker_tester

import (
	"github.com/funtimecoding/soil/pkg/chat/constant"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/mattermost/mattermost/server/public/model"
)

func ReactionEvent(
	kind model.WebsocketEventType,
	r *model.Reaction,
) *model.WebSocketEvent {
	v := &model.WebSocketEvent{}
	v = v.SetEvent(kind)

	return v.SetData(
		map[string]any{
			constant.MattermostReactionField: notation.Encode(r, false),
		},
	)
}
