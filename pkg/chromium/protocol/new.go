package protocol

import (
	"github.com/funtimecoding/soil/pkg/chromium/constant"
	"github.com/funtimecoding/soil/pkg/chromium/face"
)

func New(
	c face.ChromiumClient,
	tab string,
) *Protocol {
	t := c.MustHostTab(tab)

	if t == nil {
		panic(constant.TabNotFound)
	}

	return &Protocol{client: c, identifier: t.Identifier}
}
