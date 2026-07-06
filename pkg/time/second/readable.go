package second

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/floats"
	"github.com/funtimecoding/soil/pkg/math/round"
	"github.com/funtimecoding/soil/pkg/strings/join/key_value"
)

func Readable(seconds int) string {
	var result string
	var value float64
	var unit string

	if seconds > (60*60*24*7*4)-1 {
		value = float64(seconds) / 60 / 60 / 24 / 7 / 4
		unit = "month"
	} else if seconds > (60*60*24*7)-1 {
		value = float64(seconds) / 60 / 60 / 24 / 7
		unit = "week"
	} else if seconds > (60*60*24)-1 {
		value = float64(seconds) / 60 / 60 / 24
		unit = "day"
	} else if seconds > (60*60)-1 {
		value = float64(seconds) / 60 / 60
		unit = "hour"
	} else if seconds > 60-1 {
		value = float64(seconds) / 60
		unit = "minute"
	} else {
		value = float64(seconds)
		unit = "second"
	}

	rounded := round.Round(value, 1)

	if rounded == float64(int(value)) {
		result = fmt.Sprintf("%d %s", int(value), unit)
	} else {
		result = key_value.Space(floats.ToStringRounded(value), unit)
	}

	if rounded > 1 {
		result = fmt.Sprintf("%ss", result)
	}

	if result == "0 second" {
		result = "none"
	}

	return result
}
