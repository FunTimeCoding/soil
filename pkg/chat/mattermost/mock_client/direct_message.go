package mock_client

import "github.com/mattermost/mattermost/server/public/model"

func (c *Client) DirectMessage(
	u *model.User,
	message string,
) (*model.Post, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	result := &model.Post{
		ChannelId: model.GetDMNameFromIds(c.me.Id, u.Id),
		UserId:    c.me.Id,
		Message:   message,
	}
	c.Posted = append(c.Posted, result)

	return result, nil
}
