package mock_client

import (
	"github.com/funtimecoding/soil/pkg/kubernetes/types/native/lease"
	"github.com/funtimecoding/soil/pkg/kubernetes/types/native/pod"
	"sync"
)

type Client struct {
	Failure  error
	PodList  []*pod.Pod
	lease    *lease.Lease
	pods     map[string]bool
	podError error
	lookups  []string
	mutex    sync.Mutex
}
