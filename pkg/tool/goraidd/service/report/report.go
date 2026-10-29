package report

import "time"

type Report struct {
	Name     string
	Size     int64
	Modified time.Time
}
