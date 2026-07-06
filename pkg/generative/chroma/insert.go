package chroma

import (
	"github.com/amikos-tech/chroma-go/pkg/api/v2"
	"github.com/funtimecoding/soil/pkg/integers"
)

func (c *Client) Insert(
	l v2.Collection,
	index int,
	text string,
) {
	c.Add(
		l,
		v2.WithIDs(v2.DocumentID(integers.ToString(index))),
		v2.WithTexts(text),
	)
}
