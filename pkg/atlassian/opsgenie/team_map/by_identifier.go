package team_map

import "github.com/funtimecoding/soil/pkg/atlassian/opsgenie/team"

func (m *Map) ByIdentifier(identifier string) *team.Team {
	if v, okay := m.TeamMap[identifier]; okay {
		return v
	}

	return nil
}
