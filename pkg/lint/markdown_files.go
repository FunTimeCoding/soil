package lint

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/lint/option"
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"strings"
)

func markdownFiles(
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
				console.Format("Skip markdown file: %s\n", p)
			}

			continue
		}

		if !strings.HasSuffix(p, constant.MarkdownExtension) {
			continue
		}

		if o.Verbose {
			console.Format("Select markdown file: %s\n", p)
		}

		result = append(result, p)
	}

	return result
}
