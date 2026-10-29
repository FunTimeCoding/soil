package operation

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/typed/batch/v1"
)

func Job(
	c kubernetes.Interface,
	namespace string,
) v1.JobInterface {
	return c.BatchV1().Jobs(namespace)
}
