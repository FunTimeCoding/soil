package watcher

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost/post"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/digest/event"
	"github.com/mattermost/mattermost/server/public/model"
	"time"
)

func (w *Watcher) recordPost(v *model.WebSocketEvent) {
	p := post.Decode(v)

	if p.UserId == w.selfIdentifier() {
		return
	}

	root := p.RootId

	if root == "" {
		root = p.Id
	}

	if !w.watching(root) {
		return
	}

	w.remember(p.Id, root)
	w.Record(
		root,
		event.New(
			constant.MessageEvent,
			w.author(p.UserId),
			p.Message,
			time.UnixMilli(p.CreateAt),
		),
	)
}
