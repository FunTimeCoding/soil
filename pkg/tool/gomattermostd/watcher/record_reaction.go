package watcher

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/chat/mattermost/reaction"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/constant"
	watchConstant "github.com/funtimecoding/soil/pkg/watch/constant"
	"github.com/funtimecoding/soil/pkg/watch/digest/event"
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

	verb := "reacted"

	if kind == constant.ReactionRemovedEvent {
		verb = "removed"
	}

	author := w.author(r.UserId)
	w.Record(
		root,
		event.New(
			watchConstant.ReactionNoun,
			author,
			fmt.Sprintf("%s %s :%s:", author, verb, r.EmojiName),
			time.Now(),
		),
	)
}
