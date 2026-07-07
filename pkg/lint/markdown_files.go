package lint

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/lint/option"
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"strings"
)

func markdownFiles(
	v *virtual_file_system.System,
	skip *option.Lint,
	verbose bool,
) []string {
	var result []string

	for _, p := range v.Files() {
		if Skipped(skip, p) {
			if verbose {
				fmt.Printf("Skip markdown file: %s\n", p)
			}

			continue
		}

		if !strings.HasSuffix(p, constant.MarkdownExtension) {
			continue
		}

		if verbose {
			fmt.Printf("Select markdown file: %s\n", p)
		}

		result = append(result, p)
	}

	return result
}
