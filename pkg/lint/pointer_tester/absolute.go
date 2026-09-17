package pointer_tester

import (
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/constant"
)

func Absolute(
	path string,
	line string,
) []*concern.Concern {
	return []*concern.Concern{
		concern.NewLine(
			constant.AbsolutePointerKey,
			constant.AbsolutePointerText,
			path,
			1,
			line,
			false,
		),
	}
}
