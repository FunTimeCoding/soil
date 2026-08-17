package lease

import "k8s.io/api/coordination/v1"

func New(
	v *v1.Lease,
	cluster string,
) *Lease {
	return &Lease{Name: v.Name, Cluster: cluster, Raw: v}
}
