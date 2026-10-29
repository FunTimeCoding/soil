package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/kubernetes/client/refusal"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/cluster"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
)

func deleteResourceOn(
	x context.Context,
	c *cluster.Cluster,
	resourceType string,
	name string,
	namespace string,
) error {
	gvr, namespaced, f := c.ResolveResource(resourceType)

	if f != nil {
		return f
	}

	if namespace == "" {
		namespace = "default"
	}

	if namespaced {
		return refusal.Classify(
			c.Dynamic().Resource(gvr).Namespace(namespace).Delete(
				x,
				name,
				v1.DeleteOptions{},
			),
		)
	}

	return refusal.Classify(
		c.Dynamic().Resource(gvr).Delete(x, name, v1.DeleteOptions{}),
	)
}
