package schedule

import (
	"github.com/funtimecoding/soil/pkg/atlassian/opsgenie/team"
	"github.com/opsgenie/opsgenie-go-sdk-v2/schedule"
)

func New(
	v *schedule.Schedule,
	t *team.Team,
) *Schedule {
	return &Schedule{
		Identifier:  v.Id,
		Name:        v.Name,
		Description: v.Description,
		Team:        t,
		Raw:         v,
	}
}
