package mock_client

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/generated/client"
	"sync"
)

type Client struct {
	mutex   sync.Mutex
	Census  client.Census
	Check   client.CheckResponse
	Diff    client.DiffResponse
	Removed []string
	Asked   []string
	Failure error
	Probe   error
	Closed  int
}
