package utilization

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/generative/constant"
)

func Parse(body []byte) *Result {
	var parsed response

	if e := json.Unmarshal(body, &parsed); e != nil {
		return nil
	}

	result := &Result{}
	found := false

	for _, l := range parsed.Limits {
		switch l.Kind {
		case "session":
			result.SessionPercent = l.Percent
			result.SessionReset = resetTime(l.ResetsAt)
			found = true
		case constant.AnthropicLimitWeekly:
			result.WeeklyPercent = l.Percent
			result.WeeklyReset = resetTime(l.ResetsAt)
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
