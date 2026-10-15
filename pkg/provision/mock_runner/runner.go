package mock_runner

import (
	"github.com/funtimecoding/soil/pkg/provision/types/trigger"
	"sync"
)

type Runner struct {
	requests []trigger.Request
	fail     error
	mutex    sync.Mutex
}
