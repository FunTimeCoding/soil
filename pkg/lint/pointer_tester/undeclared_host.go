package pointer_tester

import (
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/constant"
)

func UndeclaredHost(
	path string,
	number int,
	line string,
	count int,
) []*concern.Concern {
	var result []*concern.Concern

	for range count {
		result = append(
			result,
			concern.NewLine(
				constant.UndeclaredHostKey,
				constant.UndeclaredHostText,
				path,
				number,
				line,
				false,
			),
		)
	}

	return result
}
