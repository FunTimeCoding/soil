package teams

import (
	"github.com/funtimecoding/soil/pkg/chat/teams/destination"
	"github.com/funtimecoding/soil/pkg/chat/teams/member"
	"github.com/funtimecoding/soil/pkg/chat/teams/response"
)

func (c *Client) Members(d *destination.Destination) ([]*member.Member, error) {
	items, e := c.graph.Items(d.MembersPath(), nil, 0)

	if e != nil {
		return nil, e
	}

	decoded, e := decodeAll[response.Member](items)

	if e != nil {
		return nil, e
	}

	result := make([]*member.Member, 0, len(decoded))

	for _, m := range decoded {
		result = append(result, member.New(m.DisplayName, m.User))
	}

	return result, nil
}
