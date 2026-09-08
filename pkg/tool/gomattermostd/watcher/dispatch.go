package watcher

import (
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/constant"
	"github.com/mattermost/mattermost/server/public/model"
)

func (w *Watcher) Dispatch(v *model.WebSocketEvent) {
	switch v.EventType() {
	case model.WebsocketEventPosted:
		w.recordPost(v)
	case model.WebsocketEventReactionAdded:
		w.recordReaction(v, constant.ReactionAddedEvent)
	case model.WebsocketEventReactionRemoved:
		w.recordReaction(v, constant.ReactionRemovedEvent)
	default:
	}
}
