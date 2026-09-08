package runner

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/provision/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"os"
	"time"
)

func (r *Runner) quarantine() {
	target := join.Empty(
		r.clonePath,
		constant.RunnerQuarantineSuffix,
		time.Now().UTC().Format(constant.RunnerQuarantineFormat),
	)
	errors.PanicOnError(os.Rename(r.clonePath, target))
	r.reporter.CaptureException(
		fmt.Errorf("repository quarantined: %s", target),
	)
}
