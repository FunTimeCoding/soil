package lint

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/lint/option"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"strings"
)

func runCheckers(
	v *virtual_file_system.System,
	fixes *virtual_file_system.System,
	paths []string,
	checkers []Checker,
	o *option.Lint,
	r *output.Results,
) {
	for _, p := range paths {
		if o.Verbose {
			console.Format("Process: %s\n", p)
		}

		original := v.ReadString(p)
		content := original

		for _, check := range checkers {
			result := check(p, strings.NewReader(content))

			for _, c := range result.Concerns {
				if c.Fixed && o.Fix {
					r.AddConcern(c)
				} else if !c.Fixed {
					r.AddConcern(c)
				}
			}

			if result.Fixed != "" && result.Fixed != result.Original {
				content = result.Fixed
			}
		}

		if content != original {
			fixes.WriteString(p, content)
		}
	}
}
