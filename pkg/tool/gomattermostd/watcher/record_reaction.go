package watcher

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost/reaction"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/digest/event"
	"github.com/mattermost/mattermost/server/public/model"
	"time"
)

func (w *Watcher) recordReaction(
	v *model.WebSocketEvent,
	kind constant.EventKind,
) {
	r := reaction.Decode(v)

	if r.UserId == w.selfIdentifier() {
		return
	}

	root, okay := w.rootOf(r.PostId)

	if !okay {
		return
	}

	w.Record(
		root,
		event.New(
			kind,
			w.author(r.UserId),
			join.Empty(":", r.EmojiName, ":"),
			time.Now(),
		),
	)
}
