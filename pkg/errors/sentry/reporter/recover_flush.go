package reporter

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/sentry"
	"os"
)

func (r *Reporter) RecoverFlush(v any) {
	if r.hub != nil {
		r.hub.Recover(v)
		sentry.Flush(r.hub)
	}

	if v != nil {
		fmt.Printf("Captured panic: %v\n", v)
		os.Exit(1)
	}
}
