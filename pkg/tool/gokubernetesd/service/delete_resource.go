package service

import (
	"context"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (s *Service) DeleteResource(
	x context.Context,
	clusterName string,
	resourceType string,
	name string,
	namespace string,
) error {
	c, e := s.ClusterByName(clusterName)

	if e != nil {
		return e
	}

	gvr, namespaced, f := c.ResolveResource(resourceType)

	if f != nil {
		return f
	}

	if namespace == "" {
		namespace = "default"
	}

	if namespaced {
		return c.Dynamic().Resource(gvr).Namespace(namespace).Delete(
			x,
			name,
			v1.DeleteOptions{},
		)
	}

	return c.Dynamic().Resource(gvr).Delete(x, name, v1.DeleteOptions{})
}
