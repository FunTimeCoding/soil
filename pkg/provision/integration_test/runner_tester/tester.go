package runner_tester

import (
	"github.com/funtimecoding/soil/pkg/provision/runner"
	"sync"
	"testing"
)

type Tester struct {
	t       *testing.T
	Runner  *runner.Runner
	applied []*ApplyCall
	mutex   sync.Mutex
}
