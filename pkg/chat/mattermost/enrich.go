package mattermost

import "github.com/funtimecoding/soil/pkg/chat/mattermost/post"

func (c *Client) Enrich(v []*post.Post) error {
	for _, p := range v {
		if p.User != nil {
			continue
		}

		u, e := c.User(p.UserIdentifier)

		if e != nil {
			return e
		}

		p.Enrich(u)
	}

	return nil
}
