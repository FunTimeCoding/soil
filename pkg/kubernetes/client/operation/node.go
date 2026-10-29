package operation

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/typed/core/v1"
)

func Node(c kubernetes.Interface) v1.NodeInterface {
	return c.CoreV1().Nodes()
}
