package user_map

import "github.com/funtimecoding/soil/pkg/atlassian/opsgenie/user"

func New(v []*user.User) *Map {
	m := make(map[string]*user.User, len(v))

	for _, e := range v {
		m[e.Identifier] = e
	}

	return &Map{Users: v, UserMap: m}
}
