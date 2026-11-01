package usage_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/usage_result"
)

func Summary(result *usage_result.Result) []string {
	if result == nil {
		return nil
	}

	items := []string{
		fmt.Sprintf("Session %d%%", result.FiveHourPercent),
		fmt.Sprintf("resets %s", result.FiveHourResetText()),
		fmt.Sprintf("Weekly %d%%", result.SevenDayPercent),
	}

	if result.HasFable() {
		items = append(items, fmt.Sprintf("Fable %d%%", result.FablePercent))
	}

	return items
}
