package lint

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/lint/option"
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"strings"
)

func goFiles(
	v *virtual_file_system.System,
	skip *option.Lint,
	verbose bool,
) []string {
	var result []string

	for _, p := range v.Files() {
		if Skipped(skip, p) {
			if verbose {
				console.Format("Skip go file: %s\n", p)
			}

			continue
		}

		if !strings.HasSuffix(p, constant.GoExtension) {
			continue
		}

		if IsGeneratedHeader(v.ReadString(p)) {
			if verbose {
				console.Format("Skip generated file: %s\n", p)
			}

			continue
		}

		if verbose {
			console.Format("Select go file: %s\n", p)
		}

		result = append(result, p)
	}

	return result
}
