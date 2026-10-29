package operation

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/typed/events/v1"
)

func Event(
	c kubernetes.Interface,
	namespace string,
) v1.EventInterface {
	return c.EventsV1().Events(namespace)
}
