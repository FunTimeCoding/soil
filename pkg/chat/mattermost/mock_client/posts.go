package mock_client

import "github.com/mattermost/mattermost/server/public/model"

func (c *Client) Posts(h *model.Channel) (*model.PostList, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	result := model.NewPostList()

	for _, p := range c.post {
		if p.ChannelId == h.Id {
			result.AddPost(p)
			result.AddOrder(p.Id)
		}
	}

	return result, nil
}
