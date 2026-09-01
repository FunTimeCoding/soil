package team_map

import (
	"github.com/funtimecoding/soil/pkg/atlassian/opsgenie/team"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/opsgenie/opsgenie-go-sdk-v2/alert"
)

func (m *Map) Guess(
	a *alert.Alert,
	verbose bool,
) *team.Team {
	var name string

	if m.tagToName != nil {
		name = m.tagToName(a.Tags)
	}

	if name == "" {
		if verbose {
			console.Format("No team: %+v\n", a)
		}

		return nil
	}

	for _, t := range m.Teams {
		if t.Name == name {
			return t
		}
	}

	if verbose {
		console.Format("Guess fail: %+v\n", a)
	}

	return nil
}
