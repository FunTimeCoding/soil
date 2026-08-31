package fable_snapshot

import "time"

func New(
	percent int,
	reset string,
	createdAt time.Time,
) *Snapshot {
	return &Snapshot{Percent: percent, Reset: reset, CreatedAt: createdAt}
}
