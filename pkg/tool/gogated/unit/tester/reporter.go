package tester

import "github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"

func (o *Tester) Reporter() *memory.Memory {
	return o.server.Reporter
}
