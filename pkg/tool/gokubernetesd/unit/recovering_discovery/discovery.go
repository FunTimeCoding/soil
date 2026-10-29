package recovering_discovery

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/discovery/fake"
	"sync"
)

type Discovery struct {
	*fake.FakeDiscovery
	mutex     sync.Mutex
	failures  int
	resources []*v1.APIResourceList
}
