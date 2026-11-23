package teams

import (
	"github.com/funtimecoding/soil/pkg/chat/teams/message"
	"github.com/funtimecoding/soil/pkg/chat/teams/response"
	"net/url"
)

func (c *Client) messages(
	path string,
	v url.Values,
	limit int,
) ([]*message.Message, error) {
	items, e := c.graph.Items(path, v, limit)

	if e != nil {
		return nil, e
	}

	decoded, e := decodeAll[response.Message](items)

	if e != nil {
		return nil, e
	}

	return toMessages(decoded), nil
}
