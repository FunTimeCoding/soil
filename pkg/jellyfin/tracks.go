package jellyfin

import (
	"github.com/funtimecoding/soil/pkg/jellyfin/constant"
	"github.com/funtimecoding/soil/pkg/jellyfin/item"
	"net/url"
)

func (c *Client) Tracks(albumIdentifier string) ([]*item.Item, error) {
	v := url.Values{}
	v.Set("ParentId", albumIdentifier)
	v.Set("IncludeItemTypes", "Audio")
	v.Set("Recursive", "true")
	v.Set("Fields", constant.ItemFields)
	v.Set("SortBy", "SortName")
	v.Set("SortOrder", "Ascending")
	result, _, e := c.queryItems("/Items", v)

	return result, e
}
