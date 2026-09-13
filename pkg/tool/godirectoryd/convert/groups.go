package convert

import (
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/types/group"
)

func Groups(v []*group.Group) []*server.Group {
	result := make([]*server.Group, 0, len(v))

	for _, g := range v {
		result = append(result, Group(g))
	}

	return result
}
