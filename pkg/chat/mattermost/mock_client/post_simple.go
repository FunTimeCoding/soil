package mock_client

import "github.com/mattermost/mattermost/server/public/model"

func (c *Client) PostSimple(
	channel *model.Channel,
	message string,
) (*model.Post, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	result := &model.Post{ChannelId: channel.Name, Message: message}
	c.Posted = append(c.Posted, result)

	return result, nil
}
