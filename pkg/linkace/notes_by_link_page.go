package linkace

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/linkace/note"
	"github.com/funtimecoding/soil/pkg/linkace/page"
	"github.com/funtimecoding/soil/pkg/linkace/response"
	"strconv"
)

func (c *Client) NotesByLinkPage(
	linkIdentifier int,
	p int,
) (*page.Page[*note.Note], error) {
	var r response.NoteResponse

	if e := c.basic.Get(
		fmt.Sprintf("links/%d/notes", linkIdentifier),
		map[string]string{"page": strconv.Itoa(p)},
		&r,
	); e != nil {
		return nil, e
	}

	return &page.Page[*note.Note]{
		Items:       note.NewSlice(r.Payload),
		Total:       r.Total,
		CurrentPage: r.CurrentPage,
		LastPage:    r.LastPage,
		PerPage:     r.PerPage,
	}, nil
}
