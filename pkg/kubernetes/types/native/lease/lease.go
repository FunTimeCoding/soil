package lease

import "k8s.io/api/coordination/v1"

type Lease struct {
	Name    string
	Cluster string
	Raw     *v1.Lease
}
