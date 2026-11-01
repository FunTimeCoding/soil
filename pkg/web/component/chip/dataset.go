package chip

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"maragu.dev/gomponents"
)

func (c *Chip) Dataset(
	name string,
	value string,
) *Chip {
	c.attributes = append(
		c.attributes,
		gomponents.Attr(join.Empty("data-", name), value),
	)

	return c
}
