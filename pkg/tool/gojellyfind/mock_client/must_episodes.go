package mock_client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/jellyfin/item"
)

func (c *Client) MustEpisodes(seriesIdentifier string) []*item.Item {
	result, e := c.Episodes(seriesIdentifier)
	errors.PanicOnError(e)

	return result
}
