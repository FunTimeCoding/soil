package team_map

import "github.com/funtimecoding/soil/pkg/atlassian/opsgenie/team"

func (m *Map) ByName(s string) *team.Team {
	for _, v := range m.TeamMap {
		if v.Name == s {
			return v
		}
	}

	return nil
}
