package opsgenie

import (
	"github.com/funtimecoding/soil/pkg/atlassian/opsgenie/schedule"
	"github.com/funtimecoding/soil/pkg/errors"
	rawSchedule "github.com/opsgenie/opsgenie-go-sdk-v2/schedule"
	"sort"
)

func (c *Client) Schedules() []*schedule.Schedule {
	response, e := c.userClient.Schedule.List(
		c.context,
		&rawSchedule.ListRequest{Expand: new(true)},
	)
	errors.PanicOnError(e)
	result := schedule.NewSlice(response.Schedule, c.TeamMap())
	sort.SliceStable(
		result,
		func(
			i int,
			j int,
		) bool {
			return result[i].Name < result[j].Name
		},
	)

	return result
}
