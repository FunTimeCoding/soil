package jellyfin

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/jellyfin/constant"
	"github.com/funtimecoding/soil/pkg/jellyfin/item"
	"net/url"
)

func (c *Client) Episodes(seriesIdentifier string) ([]*item.Item, error) {
	v := url.Values{}
	v.Set("Fields", constant.ItemFields)
	result, _, e := c.queryItems(
		fmt.Sprintf("/Shows/%s/Episodes", seriesIdentifier),
		v,
	)

	return result, e
}
