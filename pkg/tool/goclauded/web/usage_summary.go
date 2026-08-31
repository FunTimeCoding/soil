package web

import "fmt"

func (s *Server) usageSummary() []string {
	result := s.service.Usage()

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
