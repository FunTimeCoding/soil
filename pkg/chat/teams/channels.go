package teams

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/chat/constant"
	"github.com/funtimecoding/soil/pkg/chat/teams/channel"
	"github.com/funtimecoding/soil/pkg/chat/teams/response"
	"net/url"
)

func (c *Client) Channels(team string) ([]*channel.Channel, error) {
	items, e := c.graph.Items(
		fmt.Sprintf(constant.TeamsChannelsPath, url.PathEscape(team)),
		nil,
		0,
	)

	if e != nil {
		return nil, e
	}

	decoded, e := decodeAll[response.Channel](items)

	if e != nil {
		return nil, e
	}

	result := make([]*channel.Channel, 0, len(decoded))

	for _, h := range decoded {
		v := channel.New(h.Identifier, h.DisplayName)
		v.Description = h.Description
		v.Membership = h.MembershipType
		v.Link = h.WebLocator
		result = append(result, v)
	}

	return result, nil
}
