package model_context

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/cruise_setting"
	"time"
)

func cruiseLine(
	v *cruise_setting.Setting,
	grace time.Duration,
) string {
	result := "cruise off"

	if grace > 0 {
		result = fmt.Sprintf("cruise after %s idle", grace)
	}

	if v.Grace == nil {
		result = fmt.Sprintf("%s (default)", result)
	}

	if v.WokenAt != nil {
		result = fmt.Sprintf(
			"%s, woken %d times, last %s",
			result,
			v.Woken,
			v.WokenAt.Format(time.RFC3339),
		)
	}

	return result
}
