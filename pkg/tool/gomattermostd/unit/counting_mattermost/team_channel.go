package counting_mattermost

import "github.com/mattermost/mattermost/server/public/model"

func (c *Client) TeamChannel(name string) (*model.Channel, error) {
	c.teamChannel.Add(1)

	return c.Client.TeamChannel(name)
}
