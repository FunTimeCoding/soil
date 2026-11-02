package identified_reporter

import "github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"

func New(identifier string) *Reporter {
	return &Reporter{Memory: memory.New(), identifier: identifier}
}
