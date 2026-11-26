package pointer_tester

import (
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/constant"
)

func DeadHinted(
	path string,
	number int,
	line string,
	message string,
) []*concern.Concern {
	return []*concern.Concern{
		concern.NewLine(
			constant.DeadPointerKey,
			message,
			path,
			number,
			line,
			false,
		),
	}
}
