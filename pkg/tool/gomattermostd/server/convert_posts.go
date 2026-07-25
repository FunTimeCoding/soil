package server

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost/post"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/generated/server"
)

func convertPosts(posts []*post.Post) []*server.Post {
	result := make([]*server.Post, len(posts))

	for i, p := range posts {
		r := &server.Post{
			Identifier: p.Raw.Id,
			Message:    p.Message,
			CreatedAt:  p.Create,
		}

		if p.User != nil {
			r.Username = p.User.Username
		}

		if p.Raw.RootId != "" {
			r.Root = new(p.Raw.RootId)
		}

		if len(p.Raw.FileIds) > 0 {
			files := []string(p.Raw.FileIds)
			r.Files = &files
		}

		result[i] = r
	}

	return result
}
