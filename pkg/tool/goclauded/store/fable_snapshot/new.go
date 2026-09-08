package fable_snapshot

import "time"

func New(
	percent int,
	reset string,
	resetAt *time.Time,
	createdAt time.Time,
) *Snapshot {
	return &Snapshot{
		Percent:   percent,
		Reset:     reset,
		ResetAt:   resetAt,
		CreatedAt: createdAt,
	}
}
