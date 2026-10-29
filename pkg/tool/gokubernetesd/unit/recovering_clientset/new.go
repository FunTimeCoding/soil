package recovering_clientset

import (
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/unit/recovering_discovery"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/discovery/fake"
	kubernetesFake "k8s.io/client-go/kubernetes/fake"
)

func New(
	failures int,
	resources []*v1.APIResourceList,
) *Clientset {
	c := kubernetesFake.NewClientset()

	return &Clientset{
		Clientset: c,
		discovery: recovering_discovery.New(
			c.Discovery().(*fake.FakeDiscovery),
			failures,
			resources,
		),
	}
}
