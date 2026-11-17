package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"testing"
)

func New(
	t *testing.T,
	o ...Option,
) *Tester {
	t.Helper()
	result := &Tester{Reporter: memory.New(), t: t}
	WithInstances("alfa")(result)

	for _, f := range o {
		f(result)
	}

	return result
}
