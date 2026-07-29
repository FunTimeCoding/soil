package chroma

import (
	"github.com/amikos-tech/chroma-go/pkg/api/v2"
	"github.com/funtimecoding/soil/pkg/generative/constant"
)

func (c *Client) ClearCollection(l v2.Collection) {
	// Deletes all, even if no name field exists
	c.Delete(l, v2.WithWhere(v2.NotEqString(constant.ChromaNameField, "")))
}
