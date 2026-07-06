package lint

import (
	"github.com/funtimecoding/soil/pkg/lint/option"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
)

func Check(
	v *virtual_file_system.System,
	skip *option.Lint,
	fix bool,
	verbose bool,
	stubTest bool,
	r *output.Results,
) *virtual_file_system.System {
	fixes := virtual_file_system.New()
	paths := goFiles(v, skip, verbose)
	runCheckers(
		v,
		fixes,
		paths,
		[]Checker{
			Import,
			Function,
			Variable,
			PackageName,
			StrayConstant,
			Spacing,
			VariableGrouping,
		},
		fix,
		verbose,
		r,
	)

	if stubTest {
		generateStubTests(v, fixes, paths, fix)
	}

	runCheckers(
		v,
		fixes,
		markupFiles(v, skip, verbose),
		[]Checker{Markup},
		fix,
		verbose,
		r,
	)

	return fixes
}
