package mattermost

import "github.com/mattermost/mattermost/server/public/model"

func (c *Client) SearchChannels(
	t *model.Team,
	query string,
) ([]*model.Channel, error) {
	result, r, e := c.client.AutocompleteChannelsForTeamForSearch(
		c.context,
		t.Id,
		query,
	)

	if e != nil {
		return nil, wrapError(r, e)
	}

	return result, nil
}
