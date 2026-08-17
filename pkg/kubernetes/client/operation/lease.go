package operation

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/typed/coordination/v1"
)

func Lease(
	c *kubernetes.Clientset,
	namespace string,
) v1.LeaseInterface {
	return c.CoordinationV1().Leases(namespace)
}
