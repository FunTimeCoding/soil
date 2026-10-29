package mock_client

import "github.com/mattermost/mattermost/server/public/model"

func (c *Client) FileInformation(_ string) (*model.FileInfo, error) {
	return nil, c.Failure
}
