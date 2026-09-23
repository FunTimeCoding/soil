package connector

import "time"

func parseMoment(value *string) *time.Time {
	if value == nil {
		return nil
	}

	parsed, e := time.Parse(time.RFC3339Nano, *value)

	if e != nil {
		return nil
	}

	return &parsed
}
