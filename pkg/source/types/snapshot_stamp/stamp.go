package snapshot_stamp

import "time"

type Stamp struct {
	Size     int64
	Modified time.Time
}
