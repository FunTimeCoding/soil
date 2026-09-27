package connector

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"net/url"
)

func (c *Client) streamLocator(
	name string,
	kinds []string,
) string {
	values := url.Values{}
	values.Set(constant.Subscriber, name)

	if len(kinds) > 0 {
		values.Set(constant.Kinds, join.Comma(kinds))
	}

	return join.Empty(c.base, constant.EventStreamPath, "?", values.Encode())
}
