package mattermost

import "github.com/mattermost/mattermost/server/public/model"

func (c *Client) FileInformation(identifier string) (*model.FileInfo, error) {
	result, r, e := c.client.GetFileInfo(c.context, identifier)

	return result, wrapError(r, e)
}
