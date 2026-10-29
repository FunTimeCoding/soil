package mock_client

import "github.com/mattermost/mattermost/server/public/model"

func (c *Client) UploadFile(
	_ *model.Channel,
	_ string,
	_ []byte,
) (*model.FileInfo, error) {
	return nil, c.Failure
}
