package recovering_clientset

import (
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/unit/recovering_discovery"
	"k8s.io/client-go/kubernetes/fake"
)

type Clientset struct {
	*fake.Clientset
	discovery *recovering_discovery.Discovery
}
