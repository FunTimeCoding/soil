package teams

import (
	"github.com/funtimecoding/soil/pkg/chat/constant"
	"github.com/funtimecoding/soil/pkg/chat/teams/chat"
	"github.com/funtimecoding/soil/pkg/chat/teams/response"
	microsoftConstant "github.com/funtimecoding/soil/pkg/microsoft/constant"
	"net/url"
)

func (c *Client) Chats(limit int) ([]*chat.Chat, error) {
	limit, top := pageSize(limit)
	items, e := c.graph.Items(
		constant.TeamsChatsPath,
		url.Values{
			microsoftConstant.TopFilter:   {top},
			microsoftConstant.OrderFilter: {constant.TeamsChatOrder},
		},
		limit,
	)

	if e != nil {
		return nil, e
	}

	decoded, e := decodeAll[response.Chat](items)

	if e != nil {
		return nil, e
	}

	result := make([]*chat.Chat, 0, len(decoded))

	for _, h := range decoded {
		v := chat.New(h.Identifier, h.ChatType)
		v.Topic = h.Topic
		v.Updated = h.LastUpdated
		v.Link = h.WebLocator
		result = append(result, v)
	}

	return result, nil
}
