package string_constant

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	"os"
	"path/filepath"
	"strings"
)

func collectFromConstantDirectory(
	result map[string][]KnownConstant,
	directory string,
	subDirectory string,
) {
	constantDirectory := filepath.Join(directory, subDirectory)
	entries, e := os.ReadDir(constantDirectory)

	if e != nil {
		return
	}

	for _, n := range entries {
		if n.IsDir() || !strings.HasSuffix(n.Name(), soil.GoExtension) {
			continue
		}

		if strings.HasSuffix(n.Name(), soil.TestSuffix) {
			continue
		}

		parseConstants(
			result,
			filepath.Join(constantDirectory, n.Name()),
			subDirectory,
		)
	}
}
