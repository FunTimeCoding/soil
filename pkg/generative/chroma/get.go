package chroma

import (
	"github.com/amikos-tech/chroma-go/pkg/api/v2"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) Get(
	o v2.Collection,
	p ...v2.CollectionGetOption,
) v2.GetResult {
	result, e := o.Get(c.context, p...)
	errors.PanicOnError(e)

	return result
}
