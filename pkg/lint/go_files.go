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
	o *option.Lint,
) []string {
	var result []string

	for _, p := range v.Files() {
		if !InScope(o, p) {
			continue
		}

		if Skipped(o, p) {
			if o.Verbose {
				console.Format("Skip go file: %s\n", p)
			}

			continue
		}

		if !strings.HasSuffix(p, constant.GoExtension) {
			continue
		}

		if IsGeneratedHeader(v.ReadString(p)) {
			if o.Verbose {
				console.Format("Skip generated file: %s\n", p)
			}

			continue
		}

		if o.Verbose {
			console.Format("Select go file: %s\n", p)
		}

		result = append(result, p)
	}

	return result
}
