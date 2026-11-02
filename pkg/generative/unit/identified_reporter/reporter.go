package identified_reporter

import "github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"

type Reporter struct {
	*memory.Memory
	identifier string
}
