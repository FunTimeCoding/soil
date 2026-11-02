package usage

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/usage_result"
	"time"
)

func (s *Service) FindUsage() (*usage_result.Result, bool, error) {
	rate, e := s.store.LatestRateSnapshot()

	if e != nil || rate == nil {
		return nil, false, e
	}

	fable, f := s.store.LatestFableSnapshot()

	if f != nil {
		return nil, false, f
	}

	fablePercent := 0
	fableReset := ""
	var fableResetAt time.Time
	updated := rate.CreatedAt
	windowStart := rate.SevenDayReset.Add(-constant.SevenDayWindow)

	if fable != nil && fable.CreatedAt.After(windowStart) {
		fablePercent = fable.Percent
		fableReset = fable.Reset

		if fable.ResetAt != nil {
			fableResetAt = *fable.ResetAt
		}

		if fable.CreatedAt.After(updated) {
			updated = fable.CreatedAt
		}
	}

	return usage_result.New(
		rate.FiveHourPercent,
		rate.FiveHourReset,
		rate.SevenDayPercent,
		rate.SevenDayReset,
		fablePercent,
		fableReset,
		fableResetAt,
		updated,
	), true, nil
}
