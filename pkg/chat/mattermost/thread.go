package mattermost

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost/constant"
	"github.com/funtimecoding/soil/pkg/chat/mattermost/post"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) Thread(p *model.Post) ([]*post.Post, error) {
	if p.ReplyCount == 0 {
		return []*post.Post{}, nil
	}

	list, _, e := c.client.GetPostThread(
		c.context,
		p.Id,
		constant.EmptyEntityTag,
		false,
	)

	if e != nil {
		return nil, e
	}

	var posts []*model.Post

	for i, o := range post.FromList(list, true) {
		if i == 0 {
			continue
		}

		posts = append(posts, o)
	}

	result := post.NewSlice(posts)
	f := c.Enrich(result)

	if f != nil {
		return nil, f
	}

	return result, nil
}
