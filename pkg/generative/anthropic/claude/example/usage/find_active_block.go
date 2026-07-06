package usage

import (
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/example/common"
	"time"
)

func findActiveBlock(all []*common.Timestamped) *block {
	if len(all) == 0 {
		return nil
	}

	window := 5 * time.Hour
	n := now()
	var current *block

	for _, ts := range all {
		needNew := current == nil

		if !needNew {
			pastEnd := ts.Time.Equal(current.end) ||
				ts.Time.After(current.end)
			gap := len(current.entries) > 0 &&
				ts.Time.Sub(
					current.entries[len(current.entries)-1].Time,
				) >= window
			needNew = pastEnd || gap
		}

		if needNew {
			start := ts.Time.Truncate(time.Hour)
			current = &block{
				start: start,
				end:   start.Add(window),
			}
		}

		current.entries = append(current.entries, ts)
	}

	if current != nil && current.end.After(n) {
		return current
	}

	return nil
}
