package string_constant

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"path/filepath"
	"slices"
)

func ListedConstantFiles(
	directory string,
	list func(string) []string,
) []string {
	var result []string
	own := filepath.Join(directory, "constant.go")

	if slices.Contains(list(directory), own) {
		result = append(result, own)
	}

	result = append(result, list(filepath.Join(directory, "constant"))...)
	current := filepath.Dir(directory)

	for {
		result = append(result, list(filepath.Join(current, "constant"))...)

		if filepath.Base(current) == constant.PackageDirectory {
			break
		}

		parent := filepath.Dir(current)

		if parent == current {
			break
		}

		current = parent
	}

	return result
}
