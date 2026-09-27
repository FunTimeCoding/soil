package goclaude

import "time"

func fableResetEpoch(s *string) *int64 {
	if s == nil {
		return nil
	}

	t, e := time.Parse(time.RFC3339Nano, *s)

	if e != nil {
		return nil
	}

	return new(t.Unix())
}
