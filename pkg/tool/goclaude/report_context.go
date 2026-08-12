package goclaude

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
)

func reportContext(
	c *client.ClientWithResponses,
	input *statusLineInput,
) {
	if input.SessionID == "" {
		return
	}

	body := client.PostSessionContextJSONRequestBody{
		UsedPercentage: int(input.ContextWindow.UsedPercentage),
	}

	if input.ContextWindow.WindowSize > 0 {
		body.WindowSize = new(input.ContextWindow.WindowSize)
	}

	if input.Model.DisplayName != "" {
		body.Model = new(input.Model.DisplayName)
	}

	if input.RateLimits != nil {
		body.FiveHourPercent = new(
			int(input.RateLimits.FiveHour.UsedPercentage),
		)
		body.SevenDayPercent = new(
			int(input.RateLimits.SevenDay.UsedPercentage),
		)
		body.FiveHourReset = new(input.RateLimits.FiveHour.ResetsAt)
		body.SevenDayReset = new(input.RateLimits.SevenDay.ResetsAt)
	}

	if _, e := c.PostSessionContextWithResponse(
		context.Background(),
		input.SessionID,
		body,
	); e != nil {
		return
	}
}
