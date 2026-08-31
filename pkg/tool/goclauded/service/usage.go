package service

import "github.com/funtimecoding/soil/pkg/tool/goclauded/usage_result"

func (s *Service) Usage() *usage_result.Result {
	rate, e := s.store.LatestRateSnapshot()

	if e != nil || rate == nil {
		return nil
	}

	fablePercent := 0
	fableReset := ""
	updated := rate.CreatedAt
	fable, f := s.store.LatestFableSnapshot()

	if f == nil && fable != nil {
		fablePercent = fable.Percent
		fableReset = fable.Reset

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
		updated,
	)
}
