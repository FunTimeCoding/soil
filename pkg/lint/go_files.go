package lint

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/lint/option"
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"github.com/funtimecoding/soil/pkg/system/writer"
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

		if Skipped(o.Skips, p) {
			if o.Verbose {
				writer.Print(o.Trace, "Skip go file: %s\n", p)
			}

			continue
		}

		if !strings.HasSuffix(p, soil.GoExtension) {
			continue
		}

		if IsGeneratedHeader(v.ReadString(p)) {
			if o.Verbose {
				writer.Print(o.Trace, "Skip generated file: %s\n", p)
			}

			continue
		}

		if o.Verbose {
			writer.Print(o.Trace, "Select go file: %s\n", p)
		}

		result = append(result, p)
	}

	return result
}
