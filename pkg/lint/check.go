package lint

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/option"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/lint/repository"
	"github.com/funtimecoding/soil/pkg/lint/spacing"
	"github.com/funtimecoding/soil/pkg/lint/types/checker_step"
	"github.com/funtimecoding/soil/pkg/source/layer"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func Check(
	repo *repository.Repository,
	o *option.Lint,
	r *output.Results,
) *layer.Layer {
	fixes := o.Layer

	if fixes == nil {
		fixes = layer.New(nil)
	}

	runCheckers(
		repo,
		fixes,
		goFiles(repo.Files, o),
		[]*checker_step.Step{
			checker_step.New(constant.StepImport, Import, true),
			checker_step.New(constant.StepFunction, Function, true),
			checker_step.New(constant.StepVariable, Variable, true),
			checker_step.New(constant.StepPackageName, PackageName, true),
			checker_step.New(constant.StepStrayConstant, StrayConstant, true),
			checker_step.New(
				constant.StepFixtureDirective,
				FixtureDirective,
				true,
			),
			checker_step.New(constant.StepSpacing, spacing.Check, true),
			checker_step.New(
				constant.StepVariableGrouping,
				VariableGrouping,
				true,
			),
		},
		o,
		r,
	)
	runCheckers(
		repo,
		fixes,
		markupFiles(repo.Files, o),
		[]*checker_step.Step{
			checker_step.New(constant.StepMarkup, Markup, true),
		},
		o,
		r,
	)
	runCheckers(
		repo,
		fixes,
		markdownFiles(repo.Files, o),
		[]*checker_step.Step{
			checker_step.New(
				constant.StepPointers,
				pointerChecker(repo, o, r),
				false,
			),
			checker_step.New(
				join.Comma(
					append([]string{constant.StepReflow}, o.ReflowSkips...),
				),
				Unless(o.ReflowSkips, Reflow),
				true,
			),
			checker_step.New(
				constant.StepInterruptingList,
				InterruptingList,
				true,
			),
		},
		o,
		r,
	)

	return fixes
}
