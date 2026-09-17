package lint

import (
	"github.com/funtimecoding/soil/pkg/lint/option"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/lint/repository"
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
)

func Check(
	repo *repository.Repository,
	o *option.Lint,
	r *output.Results,
) *virtual_file_system.System {
	fixes := virtual_file_system.New()
	runCheckers(
		repo.Files,
		fixes,
		goFiles(repo.Files, o),
		[]Checker{
			Import,
			Function,
			Variable,
			PackageName,
			StrayConstant,
			FixtureDirective,
			Spacing,
			VariableGrouping,
		},
		o,
		r,
	)
	runCheckers(
		repo.Files,
		fixes,
		markupFiles(repo.Files, o),
		[]Checker{Markup},
		o,
		r,
	)
	runCheckers(
		repo.Files,
		fixes,
		markdownFiles(repo.Files, o),
		[]Checker{Pointers(resolver(repo, o), r.AddUnchecked)},
		o,
		r,
	)

	return fixes
}
