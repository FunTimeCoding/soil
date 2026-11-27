package teams

import (
	"github.com/funtimecoding/soil/pkg/chat/constant"
	"github.com/funtimecoding/soil/pkg/chat/teams/chat"
	"github.com/funtimecoding/soil/pkg/chat/teams/outgoing"
	"github.com/funtimecoding/soil/pkg/chat/teams/response"
)

func (c *Client) CreateChat(
	kind string,
	users []string,
	topic string,
) (*chat.Chat, error) {
	var created response.Chat

	if e := c.graph.Post(
		constant.TeamsCreateChatPath,
		outgoing.NewChat(kind, users, topic),
		&created,
	); e != nil {
		return nil, e
	}

	result := chat.New(created.Identifier, created.ChatType)
	result.Topic = created.Topic
	result.Updated = created.LastUpdated
	result.Link = created.WebLocator

	return result, nil
}
