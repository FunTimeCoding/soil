package convert

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/finding_group"
)

func FindingGroups(v []*finding_group.Group) []*server.FindingGroup {
	result := make([]*server.FindingGroup, len(v))

	for i, g := range v {
		result[i] = &server.FindingGroup{
			Vulnerability: g.Vulnerability,
			Systems:       g.Systems,
			NoFix:         g.NoFix,
			PaidOnly:      g.PaidOnly,
			Upgradable:    g.Upgradable,
		}
	}

	return result
}
