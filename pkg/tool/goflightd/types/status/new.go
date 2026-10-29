package status

import "time"

func New(
	events int64,
	snapshots int64,
	marks int64,
	lastEvent *time.Time,
) *Status {
	return &Status{
		Events:    events,
		Snapshots: snapshots,
		Marks:     marks,
		LastEvent: lastEvent,
	}
}
