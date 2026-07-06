package thread

import "github.com/funtimecoding/soil/pkg/chat/mattermost/post"

func New(
	root *post.Post,
	v []*post.Post,
) *Thread {
	return &Thread{Root: root, Posts: v}
}
