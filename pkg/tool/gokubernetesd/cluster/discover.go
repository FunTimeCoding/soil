package cluster

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/kubernetes/client/refusal"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *Cluster) discover() ([]*v1.APIResourceList, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.resources != nil {
		return c.resources, nil
	}

	resources, e := c.clientset.Discovery().ServerPreferredResources()

	if resources == nil {
		return nil, fmt.Errorf(
			"discover API resources: %w",
			refusal.Classify(e),
		)
	}

	c.resources = resources

	return resources, nil
}
