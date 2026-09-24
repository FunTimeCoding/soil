package crap

import "github.com/funtimecoding/soil/pkg/crap/constant"

func mark(
	score float64,
	threshold float64,
) string {
	switch {
	case score > threshold:
		return constant.MarkAbove
	case score > threshold/2:
		return constant.MarkWarn
	default:
		return constant.MarkOkay
	}
}
