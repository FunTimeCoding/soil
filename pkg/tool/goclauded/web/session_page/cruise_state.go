package session_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/cruise_setting"
	"time"
)

func cruiseState(
	setting *cruise_setting.Setting,
	grace time.Duration,
) string {
	result := "off"

	if grace > 0 {
		result = fmt.Sprintf("after %d minutes idle", int(grace.Minutes()))
	}

	if setting.Grace == nil {
		result = fmt.Sprintf("%s, the default", result)
	}

	if setting.WokenAt != nil {
		result = fmt.Sprintf(
			"%s - woken %d times, last %s",
			result,
			setting.Woken,
			woken(setting.WokenAt),
		)
	}

	return result
}
