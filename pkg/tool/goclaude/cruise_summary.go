package goclaude

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"time"
)

func CruiseSummary(r *client.CruiseResponse) string {
	result := fmt.Sprintf("cruise %s", r.Mode)

	if r.PaceMinutes != nil {
		result = fmt.Sprintf("%s every %dm", result, *r.PaceMinutes)
	}

	if r.PromotedAt != nil {
		result = fmt.Sprintf(
			"%s, woken %d times, last %s",
			result,
			r.Woken,
			r.PromotedAt.Format(time.RFC3339),
		)
	}

	return result
}
