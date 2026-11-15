package mock_client

import (
	"github.com/funtimecoding/soil/pkg/chat/constant"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
)

func (c *Client) File(identifier string) ([]byte, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	return nil, not_found.New(constant.MattermostFileField, identifier)
}
