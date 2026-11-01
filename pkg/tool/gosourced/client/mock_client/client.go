package mock_client

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/generated/client"
	"sync"
)

type Client struct {
	mutex   sync.Mutex
	Diff    client.DiffResponse
	Asked   []string
	Failure error
	Probe   error
	Closed  int
}
