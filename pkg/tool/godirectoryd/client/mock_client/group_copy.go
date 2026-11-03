package mock_client

import (
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/generated/client"
	"slices"
)

func groupCopy(g *client.Group) *client.Group {
	result := *g
	result.Member = slices.Clone(g.Member)

	if result.Member == nil {
		result.Member = []string{}
	}

	return &result
}
