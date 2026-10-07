package file_sum

import "time"

type Sum struct {
	Size     int64
	Modified time.Time
	Sum      [32]byte
}
