package outgoing

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/chat/constant"
	"net/url"
)

func NewChat(
	kind string,
	users []string,
	topic string,
) *Chat {
	result := &Chat{Kind: kind, Topic: topic}

	for _, u := range users {
		result.Members = append(
			result.Members,
			Member{
				Type:  constant.TeamsMemberType,
				Roles: []string{constant.TeamsOwnerRole},
				Binding: fmt.Sprintf(
					constant.TeamsUserBinding,
					url.PathEscape(u),
				),
			},
		)
	}

	return result
}
