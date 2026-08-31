package service

import "github.com/funtimecoding/soil/pkg/tool/goclauded/service/usage_window"

func (s *Service) UsageWindow() *usage_window.Window {
	latest, e := s.store.LatestRateSnapshot()

	if e != nil || latest == nil {
		return nil
	}

	end := latest.SevenDayReset
	start := end.AddDate(0, 0, -7)
	rate, f := s.store.RateSnapshotsSince(start)

	if f != nil {
		return nil
	}

	fable, g := s.store.FableSnapshotsSince(start)

	if g != nil {
		return nil
	}

	return usage_window.New(start, end, rate, fable)
}
