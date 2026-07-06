package string_constant

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"os"
	"path/filepath"
	"strings"
)

func collectFromConstantDirectory(
	result map[string]knownConstant,
	directory string,
	subDirectory string,
	distance int,
) {
	constantDirectory := filepath.Join(directory, subDirectory)
	entries, e := os.ReadDir(constantDirectory)

	if e != nil {
		return
	}

	for _, n := range entries {
		if n.IsDir() || !strings.HasSuffix(
			n.Name(),
			constant.GoExtension,
		) {
			continue
		}

		if strings.HasSuffix(n.Name(), constant.TestSuffix) {
			continue
		}

		parseConstants(
			result,
			filepath.Join(constantDirectory, n.Name()),
			subDirectory,
			distance,
		)
	}
}
