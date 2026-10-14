package face

import (
	"github.com/funtimecoding/soil/pkg/kubernetes/types/native/lease"
	"github.com/funtimecoding/soil/pkg/kubernetes/types/native/pod"
)

type Cluster interface {
	Lease(
		namespace string,
		name string,
	) (*lease.Lease, error)
	Pod(
		namespace string,
		name string,
	) (*pod.Pod, error)
}
