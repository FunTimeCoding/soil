package operation

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/typed/networking/v1"
)

func NetworkPolicy(
	c kubernetes.Interface,
	namespace string,
) v1.NetworkPolicyInterface {
	return c.NetworkingV1().NetworkPolicies(namespace)
}
