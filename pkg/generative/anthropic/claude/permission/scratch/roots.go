package scratch

import (
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"path/filepath"
)

func Roots(
	project string,
	directories []string,
) []string {
	result := []string{filepath.Join(project, constant.ProjectScratchDirectory)}

	for _, d := range directories {
		result = append(
			result,
			filepath.Join(d, constant.ProjectScratchDirectory),
		)
	}

	return result
}
