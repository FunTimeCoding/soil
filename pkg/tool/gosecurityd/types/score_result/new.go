package score_result

import "time"

func New(
	filled bool,
	mirrored int,
	start int,
	total int,
	synced time.Time,
) *Result {
	return &Result{
		Filled:   filled,
		Mirrored: mirrored,
		Start:    start,
		Total:    total,
		Synced:   synced,
	}
}
