package ansible_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/provision/store"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/runner"
	"testing"
)

type Tester struct {
	t        *testing.T
	Runner   *runner.Runner
	Reporter *memory.Memory
	Store    *store.Store
	capture  string
}
