package teams

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/chat/constant"
	"github.com/funtimecoding/soil/pkg/chat/teams/message"
	"github.com/funtimecoding/soil/pkg/chat/teams/outgoing"
	"github.com/funtimecoding/soil/pkg/chat/teams/response"
	"net/url"
)

func (c *Client) SendChatMessage(
	chat string,
	text string,
) (*message.Message, error) {
	var sent response.Message

	if e := c.graph.Post(
		fmt.Sprintf(constant.TeamsChatMessagesPath, url.PathEscape(chat)),
		outgoing.New(text),
		&sent,
	); e != nil {
		return nil, e
	}

	result := toMessages([]response.Message{sent})

	if len(result) == 0 {
		return nil, fmt.Errorf("graph answered no message for chat %s", chat)
	}

	return result[0], nil
}
