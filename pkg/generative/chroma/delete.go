package chroma

import (
	"github.com/amikos-tech/chroma-go/pkg/api/v2"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) Delete(
	l v2.Collection,
	o ...v2.CollectionDeleteOption,
) {
	errors.PanicOnError(l.Delete(c.context, o...))
}
