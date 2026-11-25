package mock_client

import "time"

func moment(t *time.Time) time.Time {
	if t == nil {
		return time.Time{}
	}

	return *t
}
