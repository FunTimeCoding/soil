package user

import "github.com/funtimecoding/soil/pkg/atlassian/confluence/basic/response"

func New(r *response.User) *User {
	return &User{Name: r.DisplayName, Raw: r}
}
