package cluster

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"strings"
)

func (c *Cluster) ResolveResource(
	resourceType string,
) (schema.GroupVersionResource, bool, error) {
	resources, e := c.discover()

	if e != nil {
		return schema.GroupVersionResource{}, false, e
	}

	lower := strings.ToLower(resourceType)

	for _, list := range resources {
		gv, f := schema.ParseGroupVersion(list.GroupVersion)

		if f != nil {
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
		not_found.New("resource type", resourceType)
}
