package utilization

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/utilization/response"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/utilization/usage_result"
	"github.com/funtimecoding/soil/pkg/generative/constant"
)

func Parse(body []byte) *usage_result.Result {
	var parsed response.Utilization

	if e := json.Unmarshal(body, &parsed); e != nil {
		return nil
	}

	result := usage_result.New()
	found := false

	for _, l := range parsed.Limits {
		switch l.Kind {
		case "session":
			result.SessionPercent = l.Percent
			result.SessionReset = resetTime(l.ResetsAt)
			result.SessionSeen = true
		case constant.AnthropicLimitWeekly:
			result.WeeklyPercent = l.Percent
			result.WeeklyReset = resetTime(l.ResetsAt)
			found = true
		case constant.AnthropicLimitWeeklyScoped:
			if !fableScoped(l.Scope) {
				continue
			}

			result.FablePercent = l.Percent
			result.FableReset = resetTime(l.ResetsAt)
			result.FableSeen = true
		}
	}

	if !found {
		return nil
	}

	return result
}
