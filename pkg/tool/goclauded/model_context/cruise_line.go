package model_context

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/cruise_setting"
	"time"
)

func cruiseLine(v *cruise_setting.Setting) string {
	result := fmt.Sprintf("cruise %s", v.Mode)

	if v.Pace > 0 {
		result = fmt.Sprintf("%s every %s", result, v.Pace)
	}

	if v.PromotedAt != nil {
		result = fmt.Sprintf(
			"%s, last woke %s",
			result,
			v.PromotedAt.Format(time.RFC3339),
		)
	}

	return result
}
