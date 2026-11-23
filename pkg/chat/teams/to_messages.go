package teams

import (
	"github.com/funtimecoding/soil/pkg/chat/constant"
	"github.com/funtimecoding/soil/pkg/chat/teams/message"
	"github.com/funtimecoding/soil/pkg/chat/teams/response"
	microsoftConstant "github.com/funtimecoding/soil/pkg/microsoft/constant"
)

func toMessages(v []response.Message) []*message.Message {
	result := make([]*message.Message, 0, len(v))

	for _, m := range v {
		if m.MessageType != microsoftConstant.TeamsMessageKind || m.Deleted != nil {
			continue
		}

		text := m.Body.Content

		if m.Body.ContentType == constant.TeamsHypertextContent {
			text = toText(text)
		}

		v := message.New(m.Identifier, author(m.From), text)
		v.Created = m.Created
		v.Subject = m.Subject
		v.Replies = len(toMessages(m.Replies))
		v.Link = m.WebLocator
		result = append(result, v)
	}

	return result
}
