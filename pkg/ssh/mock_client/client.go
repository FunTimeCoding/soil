package mock_client

import (
	"github.com/funtimecoding/soil/pkg/failure"
	"github.com/funtimecoding/soil/pkg/ssh/command"
	"sync"
)

type Client struct {
	mutex       sync.Mutex
	failures    *failure.Failure
	files       map[string][]byte
	directories map[string]bool
	Commands    []*command.Command
	Closed      int
}
