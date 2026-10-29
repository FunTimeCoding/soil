package operation

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/typed/core/v1"
)

func Pod(
	c kubernetes.Interface,
	namespace string,
) v1.PodInterface {
	return c.CoreV1().Pods(namespace)
}
