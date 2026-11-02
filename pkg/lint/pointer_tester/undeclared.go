package pointer_tester

import (
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/pointer"
)

func Undeclared(
	path string,
	line string,
) []*concern.Concern {
	return []*concern.Concern{
		concern.NewLine(
			constant.DeadPointerKey,
			pointer.Message(
				constant.VerdictDead,
				constant.UndeclaredSymbolHint,
			),
			path,
			1,
			line,
			false,
		),
	}
}
