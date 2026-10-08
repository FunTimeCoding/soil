package goclaude

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
)

func reportContext(
	c *client.ClientWithResponses,
	input *StatusLineInput,
) {
	if input.SessionIdentifier == "" {
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

	if input.RateLimits != nil &&
		input.RateLimits.FiveHour != nil &&
		input.RateLimits.SevenDay != nil {
		body.FiveHourPercent = new(
			int(input.RateLimits.FiveHour.UsedPercentage),
		)
		body.SevenDayPercent = new(
			int(input.RateLimits.SevenDay.UsedPercentage),
		)
		body.FiveHourReset = resetEpoch(input.RateLimits.FiveHour.ResetsAt)
		body.SevenDayReset = resetEpoch(input.RateLimits.SevenDay.ResetsAt)
	}

	if f := fableScope(input.RateLimits); f != nil && f.Utilization != nil {
		body.FablePercent = new(int(*f.Utilization))
		body.FableReset = fableResetEpoch(f.ResetsAt)
	}

	if _, e := c.PostSessionContextWithResponse(
		context.Background(),
		input.SessionIdentifier,
		body,
	); e != nil {
		return
	}
}
