package recovering_discovery

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/discovery/fake"
)

func New(
	d *fake.FakeDiscovery,
	failures int,
	resources []*v1.APIResourceList,
) *Discovery {
	return &Discovery{
		FakeDiscovery: d,
		failures:      failures,
		resources:     resources,
	}
}
