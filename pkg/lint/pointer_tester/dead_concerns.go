package pointer_tester

import (
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/constant"
)

func DeadConcerns(
	path string,
	number int,
	line string,
) []*concern.Concern {
	return []*concern.Concern{
		concern.NewLine(
			constant.DeadPointerKey,
			constant.DeadPointerText,
			path,
			number,
			line,
			false,
		),
	}
}
