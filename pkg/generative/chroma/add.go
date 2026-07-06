package chroma

import (
	"github.com/amikos-tech/chroma-go/pkg/api/v2"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) Add(
	l v2.Collection,
	o ...v2.CollectionAddOption,
) {
	errors.PanicOnError(l.Add(c.context, o...))
}
