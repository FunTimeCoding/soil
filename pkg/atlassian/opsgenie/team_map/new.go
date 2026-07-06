package team_map

import "github.com/funtimecoding/soil/pkg/atlassian/opsgenie/team"

func New(v []*team.Team) *Map {
	m := make(map[string]*team.Team, len(v))

	for _, e := range v {
		m[e.Identifier] = e
	}

	return &Map{
		Teams:        v,
		KeyByNameMap: make(map[string]string),
		TeamMap:      m,
	}
}
