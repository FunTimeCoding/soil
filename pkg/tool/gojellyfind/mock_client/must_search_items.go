package mock_client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/jellyfin/item"
)

func (c *Client) MustSearchItems(
	term string,
	types []string,
	page int,
	perPage int,
) ([]*item.Item, int) {
	result, total, e := c.SearchItems(term, types, page, perPage)
	errors.PanicOnError(e)

	return result, total
}
