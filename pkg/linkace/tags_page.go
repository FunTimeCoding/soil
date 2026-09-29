package linkace

import (
	"github.com/funtimecoding/soil/pkg/linkace/page"
	"github.com/funtimecoding/soil/pkg/linkace/response"
	"github.com/funtimecoding/soil/pkg/linkace/tag"
	"strconv"
)

func (c *Client) TagsPage(p int) (*page.Page[*tag.Tag], error) {
	var r response.TagResponse

	if e := c.basic.Get(
		"tags",
		map[string]string{"page": strconv.Itoa(p)},
		&r,
	); e != nil {
		return nil, e
	}

	return &page.Page[*tag.Tag]{
		Items:       tag.NewSlice(r.Payload, c.host),
		Total:       r.Total,
		CurrentPage: r.CurrentPage,
		LastPage:    r.LastPage,
		PerPage:     r.PerPage,
	}, nil
}
