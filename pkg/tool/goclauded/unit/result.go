package unit

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/usage_result"
	"time"
)

func result(fiveHourReset time.Time) *usage_result.Result {
	return usage_result.New(
		26,
		fiveHourReset,
		20,
		time.Date(2026, 9, 2, 21, 0, 0, 0, time.Local),
		34,
		"Wed 8:59 PM",
		time.Time{},
		time.Now(),
	)
}
