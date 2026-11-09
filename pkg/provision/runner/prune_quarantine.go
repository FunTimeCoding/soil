package runner

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/provision/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func (r *Runner) pruneQuarantine() {
	prefix := join.Empty(r.clonePath, constant.RunnerQuarantineSuffix)
	paths, e := filepath.Glob(join.Empty(prefix, "*"))
	errors.PanicOnError(e)

	for _, path := range paths {
		at, f := time.Parse(
			constant.RunnerQuarantineFormat,
			strings.TrimPrefix(path, prefix),
		)

		if f != nil || time.Since(at) < constant.RunnerQuarantineRetention {
			continue
		}

		if g := os.RemoveAll(path); g != nil {
			r.reporter.CaptureException(g)

			continue
		}

		r.logger.Structured("quarantine_removed", constant.RunnerPath, path)
	}
}
