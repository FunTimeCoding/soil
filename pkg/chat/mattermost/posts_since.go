package mattermost

import (
	"github.com/funtimecoding/soil/pkg/chat/constant"
	"github.com/funtimecoding/soil/pkg/chat/mattermost/post"
	"github.com/mattermost/mattermost/server/public/model"
	"time"
)

// PostsSince fetches posts created after the given time, replies
// and attachment-only posts included, returned oldest first. Not
// to be confused with LatestPosts, which is count-based.
//
// The server caps each posts-since response at roughly a thousand
// posts, keeping the oldest - so the fetch chunks forward, moving
// since past the newest received post until a chunk comes back
// smaller than the cap. The chunk guard bounds pathological
// windows; a window that large returns its oldest posts only.
func (c *Client) PostsSince(
	h *model.Channel,
	since time.Time,
) ([]*post.Post, error) {
	var posts []*model.Post
	seen := make(map[string]bool)
	cursor := since

	for range constant.MattermostSinceChunkLimit {
		chunk, e := c.postsSinceChunk(h, cursor)

		if e != nil {
			return nil, e
		}

		var newest time.Time

		for _, v := range chunk {
			if time.UnixMilli(v.CreateAt).Before(since) {
				continue
			}

			if seen[v.Id] {
				continue
			}

			seen[v.Id] = true
			posts = append(posts, v)

			if t := time.UnixMilli(v.CreateAt); t.After(newest) {
				newest = t
			}
		}

		if len(chunk) < constant.MattermostSinceChunkThreshold ||
			newest.IsZero() {
			break
		}

		cursor = newest
	}

	result := post.NewSlice(posts)
	f := c.Enrich(result)

	if f != nil {
		return nil, f
	}

	return result, nil
}
