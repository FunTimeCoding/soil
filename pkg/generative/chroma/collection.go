package chroma

import (
	"github.com/amikos-tech/chroma-go/pkg/api/v2"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) Collection(
	name string,
	o ...v2.CreateCollectionOption,
) v2.Collection {
	result, e := c.client.GetOrCreateCollection(c.context, name, o...)
	errors.PanicOnError(e)

	return result
}
