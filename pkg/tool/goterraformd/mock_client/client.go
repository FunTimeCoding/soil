package mock_client

import (
	"github.com/funtimecoding/soil/pkg/kubernetes/types/native/lease"
	"sync"
)

type Client struct {
	lease    *lease.Lease
	pods     map[string]bool
	podError error
	lookups  []string
	mutex    sync.Mutex
}
