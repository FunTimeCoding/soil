package usage

import "github.com/funtimecoding/soil/pkg/tool/goclauded/types/usage_window"

func (s *Service) UsageWindow() *usage_window.Window {
	latest, e := s.store.LatestRateSnapshot()

	if e != nil || latest == nil {
		return nil
	}

	end := latest.SevenDayReset
	start := end.AddDate(0, 0, -7)
	rate, f := s.store.RateSnapshots(start)

	if f != nil {
		return nil
	}

	fable, g := s.store.FableSnapshots(start)

	if g != nil {
		return nil
	}

	return usage_window.New(start, end, rate, fable)
}
