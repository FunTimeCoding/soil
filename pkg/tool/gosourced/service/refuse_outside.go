package service

import (
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
)

func refuseOutside(
	directory string,
	paths []string,
) *output.Results {
	result := output.NewResultsWithDirectory(directory)

	for _, path := range paths {
		result.AddConcern(
			concern.NewFile(
				constant.ConcernValidation,
				"outside the active module - nothing written",
				path,
				false,
			),
		)
	}

	return result
}
