package pointer_tester

import (
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/constant"
)

func DeadHeading(
	path string,
	number int,
	line string,
	text string,
) []*concern.Concern {
	return []*concern.Concern{
		concern.NewLine(
			constant.DeadHeadingKey,
			text,
			path,
			number,
			line,
			false,
		),
	}
}
