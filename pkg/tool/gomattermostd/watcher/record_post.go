package watcher

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost/post"
	"github.com/funtimecoding/soil/pkg/strings/join"
	watchConstant "github.com/funtimecoding/soil/pkg/watch/constant"
	"github.com/funtimecoding/soil/pkg/watch/digest"
	"github.com/funtimecoding/soil/pkg/watch/digest/event"
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

	if !w.WatchingRoot(root) {
		return
	}

	w.remember(p.Id, root)
	author := w.author(p.UserId)
	w.Record(
		root,
		event.New(
			watchConstant.MessageNoun,
			author,
			join.Empty(author, ": ", digest.Excerpt(p.Message)),
			time.UnixMilli(p.CreateAt),
		),
	)
}
