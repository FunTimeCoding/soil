package usage_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/usage_result"
)

func Summary(result *usage_result.Result) []string {
	if result == nil {
		return nil
	}

	var items []string

	if result.HasSession() {
		items = append(
			items,
			fmt.Sprintf("Session %d%%", result.FiveHourPercent),
			fmt.Sprintf("resets %s", result.FiveHourResetText()),
		)
	}

	items = append(items, fmt.Sprintf("Weekly %d%%", result.SevenDayPercent))

	if result.HasFable() {
		items = append(items, fmt.Sprintf("Fable %d%%", result.FablePercent))
	}

	return items
}
