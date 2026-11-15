package break_result

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/test_event/test_run"
	"github.com/funtimecoding/soil/pkg/tool/gobreak/constant"
	"github.com/funtimecoding/soil/pkg/tool/gobreak/edit_span"
	"golang.org/x/tools/cover"
)

func (v *Result) Classify(
	r *test_run.Run,
	spans []*edit_span.Span,
	profiles []*cover.Profile,
) {
	if len(r.BuildOutput()) > 0 {
		v.Status = constant.StatusCompileError
		v.Detail = r.CompilerLine()

		return
	}

	if panicked := r.Panicked(); len(panicked) > 0 {
		v.Status = constant.StatusPanic
		v.Tests = panicked

		return
	}

	if failed := r.Failed(); len(failed) > 0 {
		v.Status = constant.StatusKilled
		v.Tests = failed

		return
	}

	if packages := r.FailedPackages(); len(packages) > 0 {
		v.Status = constant.StatusRuntimeError
		v.Detail = join.Comma(packages)

		return
	}

	if never := uncovered(spans, profiles); never != nil {
		v.Status = constant.StatusNoCoverage
		v.Detail = fmt.Sprintf(
			constant.NoCoverageDetail,
			never.File,
			never.First,
		)

		return
	}

	v.Status = constant.StatusSurvived
}
