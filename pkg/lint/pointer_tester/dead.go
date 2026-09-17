package pointer_tester

import (
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/constant"
)

func Dead(
	path string,
	line string,
	count int,
) []*concern.Concern {
	var result []*concern.Concern

	for range count {
		result = append(
			result,
			concern.NewLine(
				constant.DeadPointerKey,
				constant.DeadPointerText,
				path,
				1,
				line,
				false,
			),
		)
	}

	return result
}
