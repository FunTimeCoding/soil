package pointer_tester

import (
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/constant"
)

func UnregisteredMedia(
	path string,
	line string,
	count int,
) []*concern.Concern {
	var result []*concern.Concern

	for range count {
		result = append(
			result,
			concern.NewLine(
				constant.UnregisteredMediaKey,
				constant.UnregisteredMediaText,
				path,
				1,
				line,
				false,
			),
		)
	}

	return result
}
