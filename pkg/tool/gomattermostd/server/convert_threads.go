package server

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost/post"
	"github.com/funtimecoding/soil/pkg/chat/mattermost/thread"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/generated/server"
)

func convertThreads(threads []*thread.Thread) []server.Thread {
	result := make([]server.Thread, len(threads))

	for i, t := range threads {
		root := convertPosts([]*post.Post{t.Root})[0]
		replies := make([]server.Post, len(t.Posts))

		for j, r := range convertPosts(t.Posts) {
			replies[j] = *r
		}

		result[i] = server.Thread{Root: *root, Replies: replies}
	}

	return result
}
