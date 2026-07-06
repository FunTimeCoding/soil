package opsgenie

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/opsgenie/opsgenie-go-sdk-v2/schedule"
)

func (c *Client) OnCalls(scheduleName string) *schedule.GetOnCallsResult {
	result, e := c.userClient.Schedule.GetOnCalls(
		c.context,
		&schedule.GetOnCallsRequest{
			ScheduleIdentifier: scheduleName,
			Flat:               new(true),
		},
	)
	errors.PanicOnError(e)

	return result
}
