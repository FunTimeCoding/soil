package jellyfin

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/jellyfin/constant"
	"github.com/funtimecoding/soil/pkg/jellyfin/item"
	"net/url"
)

func (c *Client) Item(identifier string) (*item.Item, error) {
	v := url.Values{}
	v.Set("Ids", identifier)
	v.Set("Fields", constant.ItemFields)
	result, _, e := queryItems(c.basic, "/Items", v)

	if e != nil {
		return nil, e
	}

	for _, i := range result {
		if sameIdentifier(i.Identifier, identifier) {
			return i, nil
		}
	}

	return nil, not_found.New(constant.ItemKind, identifier)
}
