package convert

import (
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/types/group"
)

func Group(g *group.Group) *server.Group {
	member := g.Member

	if member == nil {
		member = []string{}
	}

	return &server.Group{
		Name:              g.Name,
		Number:            g.Number,
		Member:            member,
		DistinguishedName: g.DistinguishedName,
	}
}
