package runner_tester

import (
	"github.com/funtimecoding/soil/pkg/provision/runner"
	"github.com/funtimecoding/soil/pkg/provision/types/apply_call"
	"sync"
	"testing"
)

type Tester struct {
	t         *testing.T
	Runner    *runner.Runner
	ClonePath string
	remote    string
	applied   []*apply_call.Call
	mutex     sync.Mutex
}
