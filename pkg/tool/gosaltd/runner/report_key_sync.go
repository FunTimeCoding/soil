package runner

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/classify"
)

func (r *Runner) reportKeySync() {
	e := r.keySync()

	if e == nil {
		return
	}

	if classify.Reportable(e) {
		r.reporter.Capture(context.Background(), e)

		return
	}

	r.logger.Plain("key sync: %v", e)
}
