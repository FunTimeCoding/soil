package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/resource"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/resource/filter_result"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func (s *Service) GetResource(
	x context.Context,
	clusterName string,
	q GetQuery,
) (*filter_result.FilterResult, error) {
	c, e := s.ClusterByName(clusterName)

	if e != nil {
		return nil, e
	}

	gvr, namespaced, f := c.ResolveResource(q.ResourceType)

	if f != nil {
		return nil, f
	}

	namespace := q.Namespace

	if namespace == "" {
		namespace = "default"
	}

	var result *unstructured.Unstructured
	var g error

	if namespaced {
		result, g = c.Dynamic().Resource(gvr).Namespace(namespace).Get(
			x,
			q.Name,
			v1.GetOptions{},
		)
	} else {
		result, g = c.Dynamic().Resource(gvr).Get(
			x,
			q.Name,
			v1.GetOptions{},
		)
	}

	if g != nil {
		return nil, g
	}

	return resource.FilterObject(result.Object, q.Unfiltered), nil
}
