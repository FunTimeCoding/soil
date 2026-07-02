package post

import "github.com/mattermost/mattermost/server/public/model"

func FromList(
	l *model.PostList,
	oldestFirst bool,
) []*model.Post {
	var result []*model.Post

	// Posts also contains thread-parent posts that are not part of
	// the requested page - only Order reflects the page itself.
	for _, identifier := range l.Order {
		v, found := l.Posts[identifier]

		if !found {
			continue
		}

		result = append(result, v)
	}

	if oldestFirst {
		Sort(result)
	}

	return result
}
