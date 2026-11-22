package mock_runner

import (
	"github.com/funtimecoding/soil/pkg/provision/generated/server"
	"github.com/funtimecoding/soil/pkg/provision/types/trigger"
	"sync"
)

type Runner struct {
	requests     []trigger.Request
	hostCommands []server.HostCommandRequest
	fail         error
	mutex        sync.Mutex
}
