package teams

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/chat/constant"
	"github.com/funtimecoding/soil/pkg/chat/teams/member"
	"github.com/funtimecoding/soil/pkg/chat/teams/response"
	"net/url"
)

func (c *Client) ChatMembers(chat string) ([]*member.Member, error) {
	items, e := c.graph.Items(
		fmt.Sprintf(constant.TeamsChatMembersPath, url.PathEscape(chat)),
		nil,
		0,
	)

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
