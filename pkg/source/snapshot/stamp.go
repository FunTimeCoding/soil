package snapshot

import "time"

type stamp struct {
	size     int64
	modified time.Time
}
