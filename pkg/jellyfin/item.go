package jellyfin

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/jellyfin/constant"
	"github.com/funtimecoding/soil/pkg/jellyfin/item"
	"github.com/funtimecoding/soil/pkg/jellyfin/response"
	"net/url"
)

func (c *Client) Item(identifier string) (*item.Item, error) {
	v := url.Values{}
	v.Set("Fields", constant.ItemFields)
	var r response.Item
	e := c.get(fmt.Sprintf("/Items/%s", identifier), v, &r)

	if e != nil {
		return nil, e
	}

	return toItem(&r), nil
}
