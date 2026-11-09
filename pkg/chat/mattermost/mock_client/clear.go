package mock_client

import "github.com/mattermost/mattermost/server/public/model"

func (c *Client) Clear(h *model.Channel) (int, error) {
	if c.Failure != nil {
		return 0, c.Failure
	}

	var deleted int

	for identifier, p := range c.post {
		if p.ChannelId == h.Id {
			delete(c.post, identifier)
			deleted++
		}
	}

	return deleted, nil
}
