package pointer_tester

import (
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/constant"
)

func BareSlash(
	path string,
	line string,
	count int,
) []*concern.Concern {
	var result []*concern.Concern

	for range count {
		result = append(
			result,
			concern.NewLine(
				constant.BareSlashKey,
				constant.BareSlashText,
				path,
				1,
				line,
				false,
			),
		)
	}

	return result
}
