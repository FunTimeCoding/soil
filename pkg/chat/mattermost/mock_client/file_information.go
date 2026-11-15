package mock_client

import (
	"github.com/funtimecoding/soil/pkg/chat/constant"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) FileInformation(identifier string) (*model.FileInfo, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	return nil, not_found.New(constant.MattermostFileField, identifier)
}
