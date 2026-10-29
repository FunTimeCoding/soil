package mock_client

import "github.com/mattermost/mattermost/server/public/model"

func (c *Client) PostHasReaction(
	_ *model.Post,
	_ string,
) (bool, error) {
	return false, c.Failure
}
