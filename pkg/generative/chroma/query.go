package chroma

import (
	"github.com/amikos-tech/chroma-go/pkg/api/v2"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) Query(
	l v2.Collection,
	o ...v2.CollectionQueryOption,
) v2.QueryResult {
	result, e := l.Query(c.context, o...)
	errors.PanicOnError(e)

	return result
}
