package teams

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/chat/constant"
	"github.com/funtimecoding/soil/pkg/chat/teams/member"
	"github.com/funtimecoding/soil/pkg/chat/teams/message"
	"github.com/funtimecoding/soil/pkg/chat/teams/outgoing"
	"github.com/funtimecoding/soil/pkg/chat/teams/response"
	"net/url"
)

func (c *Client) SendChatMessage(
	chat string,
	text string,
	mentioned []*member.Member,
) (*message.Message, error) {
	body, e := outgoing.New(text, mentioned)

	if e != nil {
		return nil, e
	}

	var sent response.Message

	if f := c.graph.Post(
		fmt.Sprintf(constant.TeamsChatMessagesPath, url.PathEscape(chat)),
		body,
		&sent,
	); f != nil {
		return nil, f
	}

	result := toMessages([]response.Message{sent})

	if len(result) == 0 {
		return nil, fmt.Errorf("graph answered no message for chat %s", chat)
	}

	return result[0], nil
}
