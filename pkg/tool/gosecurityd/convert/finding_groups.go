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
			Severity:      optional(g.Severity),
			Unimportant:   g.Unimportant,
			Systems:       g.Systems,
			Count:         AffectedCount(g.Count),
		}

		if g.Severity != "" {
			result[i].Score = &g.Score
		}
	}

	return result
}
