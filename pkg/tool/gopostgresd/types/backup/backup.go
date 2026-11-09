package backup

import "time"

type Backup struct {
	Instance string
	Name     string
	Time     time.Time
	Size     int64
	Checksum string
	Tables   int
}
