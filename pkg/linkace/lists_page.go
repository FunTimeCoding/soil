package linkace

import (
	"github.com/funtimecoding/soil/pkg/linkace/list"
	"github.com/funtimecoding/soil/pkg/linkace/page"
	"github.com/funtimecoding/soil/pkg/linkace/response"
	"strconv"
)

func (c *Client) ListsPage(p int) (*page.Page[*list.List], error) {
	var r response.ListResponse

	if e := c.basic.Get(
		"lists",
		map[string]string{"page": strconv.Itoa(p)},
		&r,
	); e != nil {
		return nil, e
	}

	return &page.Page[*list.List]{
		Items:       list.NewSlice(r.Payload, c.host),
		Total:       r.Total,
		CurrentPage: r.CurrentPage,
		LastPage:    r.LastPage,
		PerPage:     r.PerPage,
	}, nil
}
