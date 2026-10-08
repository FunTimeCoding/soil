package string_constant

import (
	"os"
	"path/filepath"
)

func collectFromConstantFile(
	result map[string][]KnownConstant,
	directory string,
	p string,
) {
	path := filepath.Join(directory, "constant.go")

	if _, e := os.Stat(path); e != nil {
		return
	}

	parseConstants(result, path, p)
}
