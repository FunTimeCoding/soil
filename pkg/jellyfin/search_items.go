package jellyfin

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/jellyfin/constant"
	"github.com/funtimecoding/soil/pkg/jellyfin/item"
	"net/url"
	"strings"
)

func (c *Client) SearchItems(
	term string,
	types []string,
	page int,
	perPage int,
) ([]*item.Item, int, error) {
	if page < 1 {
		page = 1
	}

	if perPage < 1 {
		perPage = 25
	}

	params := url.Values{}
	params.Set("Recursive", "true")
	params.Set("Fields", constant.ItemFields)
	params.Set("Limit", fmt.Sprint(perPage))
	params.Set("StartIndex", fmt.Sprint((page-1)*perPage))

	if term != "" {
		params.Set("SearchTerm", term)
	}

	if len(types) > 0 {
		params.Set("IncludeItemTypes", strings.Join(types, ","))
	}

	return c.queryItems("/Items", params)
}
