package mock_client

import (
	"context"
	"errors"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
)

func (c *Client) RefreshSocket(context.Context) error {
	if c.Failure != nil {
		return c.Failure
	}

	if c.Closed > 0 {
		return errors.New("mattermost client closed")
	}

	return unavailable.Format("mattermost mock: no socket")
}
