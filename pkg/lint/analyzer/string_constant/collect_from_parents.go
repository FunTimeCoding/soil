package string_constant

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"path/filepath"
)

func collectFromParents(
	result map[string][]knownConstant,
	directory string,
) {
	current := filepath.Dir(directory)

	for {
		collectFromConstantDirectory(result, current, "constant")

		if filepath.Base(current) == constant.PackageDirectory {
			break
		}

		parent := filepath.Dir(current)

		if parent == current {
			break
		}

		current = parent
	}
}
