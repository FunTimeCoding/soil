package convert

import (
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/types/user"
)

func Users(v []*user.User) []*server.User {
	result := make([]*server.User, 0, len(v))

	for _, u := range v {
		result = append(result, User(u))
	}

	return result
}
