package server

import (
	"github.com/funtimecoding/soil/pkg/time/window"
	"time"
)

func timeRange(
	start *time.Time,
	end *time.Time,
) *window.Window {
	result := window.Whole()

	if start != nil {
		result.WithStart(*start)
	}

	if end != nil {
		result.WithEnd(*end)
	}

	return result
}
