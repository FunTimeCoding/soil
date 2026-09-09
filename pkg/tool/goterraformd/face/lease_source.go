package face

import "github.com/funtimecoding/soil/pkg/kubernetes/types/native/lease"

type LeaseSource interface {
	Lease(
		namespace string,
		name string,
	) (*lease.Lease, error)
}
