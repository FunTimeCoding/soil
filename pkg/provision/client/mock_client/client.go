package mock_client

import (
	"github.com/funtimecoding/soil/pkg/failure"
	"github.com/funtimecoding/soil/pkg/provision/generated/client"
	"github.com/funtimecoding/soil/pkg/provision/types/change"
	"sync"
)

type Client struct {
	Closed            int
	address           string
	calls             int
	triggered         []*change.Change
	hostCommands      []client.HostCommandRequest
	hostCommandStatus string
	failures          *failure.Failure
	mutex             sync.Mutex
}
