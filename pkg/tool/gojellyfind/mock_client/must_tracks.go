package mock_client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/jellyfin/item"
)

func (c *Client) MustTracks(albumIdentifier string) []*item.Item {
	result, e := c.Tracks(albumIdentifier)
	errors.PanicOnError(e)

	return result
}
