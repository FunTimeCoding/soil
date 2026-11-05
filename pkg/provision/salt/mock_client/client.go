package mock_client

import (
	"github.com/funtimecoding/soil/pkg/failure"
	"github.com/funtimecoding/soil/pkg/provision/salt/basic/response"
	"sync"
)

type Client struct {
	failures   *failure.Failure
	mutex      sync.Mutex
	keys       response.KeysReturn
	highstates int
	closed     int
}
