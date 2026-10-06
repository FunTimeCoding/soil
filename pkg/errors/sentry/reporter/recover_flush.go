package reporter

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/sentry"
	"os"
)

func (r *Reporter) RecoverFlush(v any) {
	if r.hub != nil {
		r.hub.Recover(v)
		sentry.Flush(r.hub)
	}

	if v != nil {
		errors.Printf("Captured panic: %v\n", v)
		os.Exit(1)
	}
}
