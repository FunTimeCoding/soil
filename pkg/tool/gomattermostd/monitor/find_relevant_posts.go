package monitor

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost/post"
	"strings"
)

func findRelevantPosts(
	posts []*post.Post,
	topics []string,
) []*post.Post {
	var result []*post.Post

	for _, p := range posts {
		lower := strings.ToLower(p.Message)

		for _, topic := range topics {
			if strings.Contains(lower, strings.ToLower(topic)) {
				result = append(result, p)

				break
			}
		}
	}

	return result
}
