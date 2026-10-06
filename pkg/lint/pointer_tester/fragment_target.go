package pointer_tester

import (
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/constant"
)

func FragmentTarget(
	path string,
	line string,
) []*concern.Concern {
	return []*concern.Concern{
		concern.NewLine(
			constant.FragmentTargetKey,
			constant.FragmentTargetText,
			path,
			1,
			line,
			false,
		),
	}
}
