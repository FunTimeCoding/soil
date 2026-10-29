package watcher

import "github.com/funtimecoding/soil/pkg/chat/mattermost/post"

func (w *Watcher) Index(root string) {
	w.remember(root, root)
	var replies []*post.Post
	e := w.attempt(
		func() error {
			p, f := w.client.FindPost(root)

			if f != nil {
				return f
			}

			replies, f = w.client.Thread(p)

			return f
		},
	)

	if e != nil {
		w.report(e)

		return
	}

	for _, v := range replies {
		w.remember(v.Identifier, root)
	}
}
