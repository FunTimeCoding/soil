package cluster

import (
	"fmt"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"strings"
)

func (c *Cluster) ResolveResource(
	resourceType string,
) (schema.GroupVersionResource, bool, error) {
	c.discoveryOnce.Do(
		func() {
			c.resources, c.discoveryFail = c.clientset.Discovery().ServerPreferredResources()
		},
	)

	if c.resources == nil {
		return schema.GroupVersionResource{}, false,
			fmt.Errorf("failed to discover API resources: %w", c.discoveryFail)
	}

	lower := strings.ToLower(resourceType)

	for _, list := range c.resources {
		gv, e := schema.ParseGroupVersion(list.GroupVersion)

		if e != nil {
			continue
		}

		for _, r := range list.APIResources {
			if strings.ToLower(r.Name) == lower || strings.ToLower(r.Kind) == lower {
				return schema.GroupVersionResource{
					Group:    gv.Group,
					Version:  gv.Version,
					Resource: r.Name,
				}, r.Namespaced, nil
			}

			for _, short := range r.ShortNames {
				if strings.ToLower(short) == lower {
					return schema.GroupVersionResource{
						Group:    gv.Group,
						Version:  gv.Version,
						Resource: r.Name,
					}, r.Namespaced, nil
				}
			}
		}
	}

	return schema.GroupVersionResource{}, false,
		fmt.Errorf("unknown resource type: %s", resourceType)
}
