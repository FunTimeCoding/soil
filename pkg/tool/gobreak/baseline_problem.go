package gobreak

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/test_event/test_run"
	"github.com/funtimecoding/soil/pkg/tool/gobreak/constant"
)

func baselineProblem(r *test_run.Run) string {
	if output := r.BuildOutput(); len(output) > 0 {
		return fmt.Sprintf(constant.BaselineBuild, compilerLine(output))
	}

	if failed := r.Failed(); len(failed) > 0 {
		return fmt.Sprintf(constant.BaselineFailed, join.Space(failed...))
	}

	if packages := r.FailedPackages(); len(packages) > 0 {
		return fmt.Sprintf(constant.BaselineSetup, join.Comma(packages))
	}

	if len(r.Tests()) == 0 {
		return constant.BaselineNoTests
	}

	return ""
}
