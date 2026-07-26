package uints64

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/uints64/constant"
)

func Readable(bytes uint64) string {
	if bytes < constant.Unit {
		return fmt.Sprintf("%d B", bytes)
	}

	index := 0
	divisor := constant.Unit

	for i := bytes / constant.Unit; i >= constant.Unit; i /= constant.Unit {
		divisor *= constant.Unit
		index++
	}

	return fmt.Sprintf(
		"%.1f %cB",
		float64(bytes)/float64(divisor),
		constant.UnitLetter[index],
	)
}
