package user

import "github.com/funtimecoding/soil/pkg/atlassian/confluence/basic/response"

type User struct {
	Name string
	Raw  *response.User
}
