package mock_client

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/chat/teams/destination"
	"github.com/funtimecoding/soil/pkg/chat/teams/member"
	"github.com/funtimecoding/soil/pkg/chat/teams/message"
)

func (c *Client) Post(
	d *destination.Destination,
	text string,
	mentioned []*member.Member,
) (*message.Message, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	store, k, e := c.target(d)

	if e != nil {
		return nil, e
	}

	result := message.New(fmt.Sprintf("sent-%d", len(store[k])+1), "", text)
	store[k] = append(store[k], result)
	c.Mentioned = append(c.Mentioned, mentioned...)

	return result, nil
}
