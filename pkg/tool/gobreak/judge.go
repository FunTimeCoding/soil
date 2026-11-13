package gobreak

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/test_event/test_run"
	"github.com/funtimecoding/soil/pkg/tool/gobreak/break_result"
	"github.com/funtimecoding/soil/pkg/tool/gobreak/constant"
	"github.com/funtimecoding/soil/pkg/tool/gobreak/edit_span"
	"golang.org/x/tools/cover"
)

func judge(
	result *break_result.Result,
	r *test_run.Run,
	spans []*edit_span.Span,
	profiles []*cover.Profile,
) {
	if output := r.BuildOutput(); len(output) > 0 {
		result.Kind = constant.ResultBuildFailed
		result.Detail = compilerLine(output)

		return
	}

	if panicked := r.Panicked(); len(panicked) > 0 {
		result.Kind = constant.ResultPanic
		result.Tests = panicked

		return
	}

	if failed := r.Failed(); len(failed) > 0 {
		result.Kind = constant.ResultRed
		result.Tests = failed

		return
	}

	if packages := r.FailedPackages(); len(packages) > 0 {
		result.Kind = constant.ResultSetupFailed
		result.Detail = join.Comma(packages)

		return
	}

	if never := unreached(spans, profiles); never != nil {
		result.Kind = constant.ResultUnreached
		result.Detail = fmt.Sprintf(
			constant.ReachedNever,
			never.File,
			never.First,
		)

		return
	}

	result.Kind = constant.ResultSurvived
}
