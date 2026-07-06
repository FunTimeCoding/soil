package user_map

import "github.com/funtimecoding/soil/pkg/atlassian/opsgenie/user"

type Map struct {
	Users   []*user.User
	UserMap map[string]*user.User
}
