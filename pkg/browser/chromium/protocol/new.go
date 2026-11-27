package protocol

import (
	"github.com/funtimecoding/soil/pkg/browser/constant"
	"github.com/funtimecoding/soil/pkg/browser/face"
)

func New(
	c face.ChromiumClient,
	tab string,
) *Protocol {
	t := c.MustHostTab(tab)

	if t == nil {
		panic(constant.ChromiumTabNotFound)
	}

	return &Protocol{client: c, identifier: t.Identifier}
}
