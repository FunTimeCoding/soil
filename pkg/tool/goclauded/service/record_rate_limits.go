package service

import "time"

func (s *Service) RecordRateLimits(
	fiveHourPercent int,
	sevenDayPercent int,
	fiveHourReset time.Time,
	sevenDayReset time.Time,
) error {
	latest, e := s.store.LatestRateSnapshot()

	if e != nil {
		return e
	}

	if latest != nil &&
		latest.FiveHourPercent == fiveHourPercent &&
		latest.SevenDayPercent == sevenDayPercent {
		return nil
	}

	s.store.SaveRateSnapshot(
		fiveHourPercent,
		sevenDayPercent,
		fiveHourReset,
		sevenDayReset,
	)
	s.store.TrimRateSnapshots()

	return nil
}
