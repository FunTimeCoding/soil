package protocol

import (
	"github.com/funtimecoding/soil/pkg/chromium"
	"github.com/funtimecoding/soil/pkg/chromium/constant"
)

func New(
	c *chromium.Client,
	tab string,
) *Protocol {
	t := c.MustTabByHost(tab)

	if t == nil {
		panic(constant.TabNotFound)
	}

	return &Protocol{client: c, context: c.TargetContext(t.Identifier)}
}
