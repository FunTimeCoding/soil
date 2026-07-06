package schedule

import (
	"github.com/funtimecoding/soil/pkg/atlassian/opsgenie/team_map"
	"github.com/opsgenie/opsgenie-go-sdk-v2/schedule"
)

func NewSlice(
	v []schedule.Schedule,
	m *team_map.Map,
) []*Schedule {
	var result []*Schedule

	for _, e := range v {
		result = append(
			result,
			New(&e, m.ByIdentifier(e.OwnerTeam.Id)),
		)
	}

	return result
}
