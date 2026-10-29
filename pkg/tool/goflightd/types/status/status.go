package status

import "time"

type Status struct {
	Events    int64
	Snapshots int64
	Marks     int64
	LastEvent *time.Time
}
