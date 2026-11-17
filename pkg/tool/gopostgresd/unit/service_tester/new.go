package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/dump/mock_dump"
	"testing"
)

func New(
	t *testing.T,
	o ...Option,
) *Tester {
	t.Helper()
	result := &Tester{
		Reporter:  memory.New(),
		Dump:      mock_dump.New(),
		Directory: t.TempDir(),
		t:         t,
	}
	WithInstances("alfa")(result)

	for _, f := range o {
		f(result)
	}

	return result
}
