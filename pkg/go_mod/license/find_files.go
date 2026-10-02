package license

import (
	"github.com/funtimecoding/soil/pkg/go_mod/constant"
	"github.com/funtimecoding/soil/pkg/system"
	"strings"
)

func FindFiles(directory string) []string {
	var result []string

	for _, e := range system.ReadDirectory(directory) {
		if e.IsDir() {
			continue
		}

		upper := strings.ToUpper(e.Name())

		for _, n := range constant.LicenseFileNames {
			if strings.HasPrefix(upper, n) {
				result = append(result, e.Name())

				break
			}
		}
	}

	return result
}
