package message

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/refusal"
	"strconv"
	"time"
)

func ParseGrace(text string) (*time.Duration, error) {
	switch text {
	case constant.CruiseDefault:
		return nil, nil
	case constant.CruiseOff:
		return new(time.Duration(0)), nil
	}

	minutes, e := strconv.Atoi(text)

	if e != nil ||
		minutes < constant.CruiseGraceMinimumMinutes ||
		minutes > constant.CruiseGraceMaximumMinutes {
		return nil, refusal.New(
			fmt.Sprintf(
				"cruise is off, default, or a whole number of minutes from %d to %d, not %q",
				constant.CruiseGraceMinimumMinutes,
				constant.CruiseGraceMaximumMinutes,
				text,
			),
		)
	}

	return new(time.Duration(minutes) * time.Minute), nil
}
