package lint

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/lint/option"
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"strings"
)

func markupFiles(
	v *virtual_file_system.System,
	skip *option.Lint,
	verbose bool,
) []string {
	var result []string

	for _, p := range v.Files() {
		if Skipped(skip, p) {
			if verbose {
				fmt.Printf("Skip markup file: %s\n", p)
			}

			continue
		}

		if !strings.HasSuffix(p, constant.MarkupExtension) &&
			!strings.HasSuffix(p, constant.ShortMarkupExtension) {
			continue
		}

		if verbose {
			fmt.Printf("Select markup file: %s\n", p)
		}

		result = append(result, p)
	}

	return result
}
