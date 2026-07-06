package team_map

import "github.com/funtimecoding/soil/pkg/atlassian/opsgenie/team"
import rawTeam "github.com/opsgenie/opsgenie-go-sdk-v2/team"

func (m *Map) Add(
	identifier string,
	name string,
	description string,
) {
	t := team.New(
		&rawTeam.ListedTeams{
			TeamMeta:    rawTeam.TeamMeta{Id: identifier, Name: name},
			Description: description,
		},
	)
	m.Teams = append(m.Teams, t)
	m.TeamMap[identifier] = t
}
