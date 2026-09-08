package utilization

import "time"

func resetTime(s *string) time.Time {
	if s == nil {
		return time.Time{}
	}

	result, e := time.Parse(time.RFC3339Nano, *s)

	if e != nil {
		return time.Time{}
	}

	return result
}
