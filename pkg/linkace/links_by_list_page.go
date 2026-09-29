package linkace

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/linkace/link"
	"github.com/funtimecoding/soil/pkg/linkace/page"
	"github.com/funtimecoding/soil/pkg/linkace/response"
	"strconv"
)

func (c *Client) LinksByListPage(
	identifier int,
	p int,
) (*page.Page[*link.Link], error) {
	var r response.LinkResponse

	if e := c.basic.Get(
		fmt.Sprintf("lists/%d/links", identifier),
		map[string]string{"page": strconv.Itoa(p)},
		&r,
	); e != nil {
		return nil, e
	}

	return &page.Page[*link.Link]{
		Items:       link.NewSlice(r.Payload, c.host),
		Total:       r.Total,
		CurrentPage: r.CurrentPage,
		LastPage:    r.LastPage,
		PerPage:     r.PerPage,
	}, nil
}
