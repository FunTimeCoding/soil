package lint

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"path/filepath"
	"strings"
)

func missingTestDirectories(
	paths []string,
	generatedPaths []string,
) map[string]string {
	directories := make(map[string]bool)
	representative := make(map[string]string)

	for _, p := range paths {
		d := filepath.Dir(p)

		if _, seen := directories[d]; !seen {
			directories[d] = false
			representative[d] = p
		}
	}

	for _, p := range generatedPaths {
		d := filepath.Dir(p)

		if _, seen := directories[d]; !seen {
			directories[d] = false
			representative[d] = p
		}
	}

	for _, p := range paths {
		if strings.HasSuffix(p, constant.TestSuffix) {
			directories[filepath.Dir(p)] = true
		}
	}

	result := make(map[string]string)

	for d, hasTest := range directories {
		if !hasTest {
			result[d] = representative[d]
		}
	}

	return result
}
