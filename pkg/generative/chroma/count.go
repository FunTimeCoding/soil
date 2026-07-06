package chroma

import (
	"github.com/amikos-tech/chroma-go/pkg/api/v2"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) Count(l v2.Collection) int {
	result, e := l.Count(c.context)
	errors.PanicOnError(e)

	return result
}
