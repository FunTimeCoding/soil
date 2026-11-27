package teams

import (
	"github.com/funtimecoding/soil/pkg/chat/teams/response"
	"github.com/funtimecoding/soil/pkg/chat/teams/user"
)

func toUser(u *response.User) *user.User {
	result := user.New(u.Identifier, u.DisplayName, u.Mail)
	result.Principal = u.Principal
	result.Title = u.JobTitle

	return result
}
