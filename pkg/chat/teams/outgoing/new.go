package outgoing

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/chat/constant"
	"github.com/funtimecoding/soil/pkg/chat/teams/member"
	"github.com/funtimecoding/soil/pkg/chat/teams/response"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"html"
	"strings"
)

func New(
	text string,
	mentioned []*member.Member,
) (*Message, error) {
	content := strings.ReplaceAll(html.EscapeString(text), "\n", "<br>")
	result := &Message{}

	for i, m := range mentioned {
		name := html.EscapeString(m.Name)
		marker := fmt.Sprintf("@%s", name)

		if !strings.Contains(content, marker) {
			return nil, validation.New("text does not name @%s", m.Name)
		}

		content = strings.Replace(
			content,
			marker,
			fmt.Sprintf("<at id=\"%d\">%s</at>", i, name),
			1,
		)
		result.Mentions = append(
			result.Mentions,
			Mention{
				Identifier: i,
				Text:       m.Name,
				Mentioned: Mentioned{
					User: MentionedUser{
						Identifier:   m.User,
						DisplayName:  m.Name,
						IdentityType: constant.TeamsUserIdentity,
					},
				},
			},
		)
	}

	result.Body = response.Body{
		ContentType: constant.TeamsHypertextContent,
		Content:     content,
	}

	return result, nil
}
